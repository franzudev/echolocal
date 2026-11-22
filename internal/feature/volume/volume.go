package volume

import (
	"sync"

	sharedvolume "github.com/franzudev/libcountertop/pkg/audio/volume"
	sharedview "github.com/franzudev/libcountertop/pkg/display/volume"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(38),
		component.Needs(board.Panel),
	)
}

type Change = sharedvolume.Change

type Volume struct {
	Changed sharedvolume.Changes

	card *sharedview.Card
}

var (
	once   sync.Once
	shared *Volume
)

func Get() *Volume {
	once.Do(func() {
		shared = &Volume{card: sharedview.NewCard(shell.Get())}

		media.Get().Levels.Changed.Listen(func(c Change) { shared.changed(c.Stream) })
		media.Get().SetScreen(shared)
	})
	return shared
}

func (v *Volume) Selected() (config.Stream, bool) { return v.card.Selected() }

func (v *Volume) Top() any { return shell.Get().Top() }

func (v *Volume) Name() string { return "volume" }

func (v *Volume) Level(s config.Stream) int { return media.Get().Levels.Level(s) }

func (v *Volume) Set(s config.Stream, level int) {
	level = sharedvolume.Clamp(level)
	if v.Level(s) == level {
		return
	}
	media.Get().Levels.Set(s, level)
}

func (v *Volume) changed(s config.Stream) {
	v.Changed.Emit(Change{Stream: s, Level: v.Level(s)})
	if showing(s) {
		shell.Get().Redraw()
		return
	}
	v.card.Show(s)
}

func (v *Volume) Watch(stream config.Stream, changed func(int)) func() {
	return v.Changed.Watch(stream, changed)
}
