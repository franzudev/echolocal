package photos

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/media/photos"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/product"
	"github.com/franzudev/echolocal/internal/hardware/display"
)

type componentPhotos struct{ *photos.Library }

func (p componentPhotos) Restore(c config.Config) { p.Library.Restore(c.Photos) }

var get = sync.OnceValue(func() *photos.Library {
	return photos.Build(product.Get(), config.PhotosSection, func() int {
		w, h := display.Get().Native()
		return max(w, h)
	})
})

func Get() *photos.Library { return get() }

func init() {
	component.Register(
		sharedcomponent.Network,
		func() componentPhotos { return componentPhotos{Get()} },
		sharedcomponent.Order(24),
		component.Needs(board.Panel),
	)
}
