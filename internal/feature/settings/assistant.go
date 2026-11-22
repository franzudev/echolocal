//go:build !board_biscuit

package settings

import (
	"github.com/franzudev/libcountertop/pkg/assistant/turn"
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"
	sharedha "github.com/franzudev/libcountertop/pkg/homeassistant"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/arbitration"
	"github.com/franzudev/echolocal/internal/feature/detect"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/tones"
	"github.com/franzudev/echolocal/internal/feature/voicematch/panel"
	"github.com/franzudev/echolocal/internal/feature/wakeword"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
	"github.com/franzudev/echolocal/internal/layout"
	"github.com/franzudev/echolocal/internal/lib/wake"
)

func assistantPages() *sharedsettings.AssistantPages {
	mac := func() string {
		mac, _ := layout.FactoryMAC()
		return mac
	}
	return sharedsettings.NewAssistantPages(
		sharedsettings.AssistantDependencies{
			Slots:       wakeword.Slots,
			Pryon:       detect.Get().Pryon(),
			PryonStop:   detect.Get().Pryon().Stop(),
			Gemini:      detect.Get().Gemini(),
			Arbitration: &arbitration.Setting,
			Preferences: preferences(),
			Models:      wake.Lib().Ours,
			Phrases:     wake.Lib().Phrases,
			Chooser:     sharedha.WakeWords{Controller: homeassistant.Get(), MAC: mac},
			Pipelines:   sharedha.Pipelines{Controller: homeassistant.Get(), MAC: mac},
			Requested:   &wakeword.Requested,
			Words: turn.Words{
				Settings: config.WakeSection,
				Speaker:  speaker.Sound(),
				Sounds:   tones.Get(),
			},
			Look:       lookPages(),
			Shell:      shell.Get(),
			VoiceMatch: panel.Get().Page,
		},
	)
}
