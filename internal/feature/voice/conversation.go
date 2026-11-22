package voice

import (
	"context"

	"github.com/franzudev/libcountertop/pkg/assistant/turn"
	"github.com/franzudev/libcountertop/pkg/audio/sound"
	sharedwake "github.com/franzudev/libcountertop/pkg/inference/wake"
	"github.com/franzudev/libcountertop/pkg/inference/wakeslots"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/activity"
	"github.com/franzudev/echolocal/internal/feature/arbitration"
	"github.com/franzudev/echolocal/internal/feature/feedback"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/mute"
	"github.com/franzudev/echolocal/internal/feature/recording"
	"github.com/franzudev/echolocal/internal/feature/tones"
	"github.com/franzudev/echolocal/internal/feature/voicematch"
	"github.com/franzudev/echolocal/internal/hardware/mic"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
	"github.com/franzudev/echolocal/internal/lib/wake"
)

type voiceSound struct{ *speaker.Driver }

func (s voiceSound) Claim(name string, fn func(context.Context, turn.Speaker) error) turn.Claim {
	return s.Driver.Claim(
		name,
		func(ctx context.Context, p *speaker.Player) error { return fn(ctx, p) },
	)
}
func (s voiceSound) Duck(on bool) { s.Backgrounds().Duck(on) }

func dependencies() turn.Dependencies {
	return turn.Dependencies{
		Source:       mic.Get(),
		Speaker:      speaker.Get(),
		Sound:        voiceSound{speaker.Sound()},
		Player:       media.Get(),
		Log:          activity.Get(),
		Recorder:     recording.Get(),
		Models:       func() []sharedwake.Model { return wake.Lib().Ours() },
		HardwareTail: sound.HardwareTail,
		Failure:      feedback.Failure,
		Cancelled:    feedback.Cancelled,
		Words: turn.Words{
			Settings: config.WakeSection,
			Speaker:  speaker.Sound(),
			Sounds:   tones.Get(),
		},
		Muted: mute.Get().Muted,
		Presentation: &turn.Presentation{
			Settings: config.WakeSection,
			Shown:    &Shown,
			Effects:  &turnEffects{},
		},
		Arbiter: arbitration.Get(),
		Voices:  voicematch.Get(),
	}
}

func wanted(models []sharedwake.Model, count int) []string {
	saved := config.Get().Wake.Words
	var ids []string
	if saved != nil {
		ids = make([]string, len(saved))
		for i, word := range saved {
			ids[i] = word.ID
		}
	}
	installed := make([]string, len(models))
	for i, m := range models {
		installed[i] = m.ID
	}
	return wakeslots.Initial(ids, installed, sharedwake.DefaultModel, count)
}
