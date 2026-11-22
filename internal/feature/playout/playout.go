package playout

import (
	"github.com/franzudev/libcountertop/pkg/media/pcm"
	"github.com/franzudev/libcountertop/pkg/media/playout"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
)

func New() *playout.Output {
	a := speaker.Sound().Backgrounds()
	return playout.New(
		playout.Options{
			Queue:    speaker.Get().Lane(config.StreamMedia),
			Rate:     speaker.Rate,
			Channels: speaker.Channels,
			DuckDB:   func() float64 { return config.Get().Media.DuckDB },
			Media:    func() playout.Media { return media.Get() },
			Took:     func(p pcm.Producer) { a.Took(p) },
			Gave:     func(p pcm.Producer) { a.Gave(p) },
			Screen:   screen(),
		},
	)
}
