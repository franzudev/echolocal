//go:build !board_biscuit

package airplay

import (
	"image"

	sharedairplay "github.com/franzudev/libcountertop/pkg/media/airplay"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/hardware/display"
)

func screen(o *sharedairplay.Options) {
	b := component.Board()
	o.Size = func() (int, int) {
		w, h := display.Get().Native()
		return display.Get().Orientation().Size(w, h)
	}
	o.Largest = image.Pt(b.DecodeWidth, b.DecodeHeight)
}
