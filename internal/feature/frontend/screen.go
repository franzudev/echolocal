//go:build !board_biscuit

package frontend

import (
	"image"
	"net/http"

	"github.com/franzudev/libcountertop/pkg/display/desktop"
	"github.com/franzudev/libcountertop/pkg/input/textentry"
	"github.com/franzudev/libcountertop/pkg/input/touch"
	"github.com/franzudev/libcountertop/pkg/runtime/control/screencmd"

	"github.com/franzudev/echolocal/internal/hardware/display"
	hwtouch "github.com/franzudev/echolocal/internal/hardware/touch"
)

var shot = func() (image.Image, error) { return screencmd.Composed(display.Get()) }

func addDesktop(pages map[string]http.Handler, input *textentry.Controller) {
	pages["desktop"] = desktop.New(
		display.Get(),
		func(c touch.Contact) { hwtouch.Get().Feed(c) },
		input,
	)
}
