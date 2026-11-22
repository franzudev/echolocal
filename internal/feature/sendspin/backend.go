package sendspin

import (
	core "github.com/franzudev/libcountertop/pkg/audio/sendspin"
	"github.com/franzudev/libcountertop/pkg/audio/sound"
	sharedvolume "github.com/franzudev/libcountertop/pkg/audio/volume"
	sharedplayer "github.com/franzudev/libcountertop/pkg/media/sendspin"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
	"github.com/franzudev/echolocal/internal/layout"
)

const Port = core.Port

type outputSink struct{ p *speaker.Player }

func (s outputSink) Attach(r core.Renderer) { s.p.Attach(r) }
func (s outputSink) Written() uint64        { return s.p.Written() }

func newOutput() *core.Output {
	return core.NewOutput(outputSink{speaker.Get()}, core.Controls{
		Volume:       sharedvolume.For(media.Get().Levels, config.StreamMedia),
		Media:        media.Get(),
		Settings:     config.MediaSection,
		HardwareTail: sound.HardwareTail,
	})
}

type identity struct{}

func (identity) Name() string { return config.Get().Device.Name }
func (identity) Identity() core.Identity {
	mac, _ := layout.FactoryMAC()
	model := config.Get().Device.Board.Model
	return core.Identity{
		ID:           mac,
		Model:        model,
		Manufacturer: layout.Manufacturer,
		Version:      layout.Version,
		ArtworkSize:  sharedplayer.ArtworkSize,
	}
}
