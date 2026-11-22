package poster

import (
	"sync"

	sharedposter "github.com/franzudev/libcountertop/pkg/display/poster"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/photos"
	hwtouch "github.com/franzudev/echolocal/internal/hardware/touch"
)

type componentPoster struct{ *sharedposter.Poster }

func (p *componentPoster) Restore(c config.Config) { p.Poster.Restore(c.Poster) }

var get = sync.OnceValue(func() *sharedposter.Poster {
	p := sharedposter.New(
		sharedposter.Dependencies{
			Settings: config.PosterSection,
			Touch:    hwtouch.Get(),
			Photos:   photos.Get(),
			DeviceID: component.DeviceScreen,
		},
	)
	component.Settings.Add(p.Settings())
	return p
})

func Get() *sharedposter.Poster { return get() }
func init() {
	component.Register(
		sharedcomponent.Device,
		func() *componentPoster { return &componentPoster{Get()} },
		sharedcomponent.Order(25),
		component.Needs(board.Panel),
	)
}
