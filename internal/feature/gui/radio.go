//go:build !board_biscuit

package gui

import (
	gogui "github.com/go-gui-org/go-gui/gui"

	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/tuner"
	"github.com/franzudev/libcountertop/pkg/display/widgets"

	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/radio"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func extendRadio(look *widgets.UI, o *screens.Options, app func() *screens.App) {
	dial := tuner.New(tuner.Options{
		UI:     look,
		Shell:  shell.Get(),
		Radio:  radio.Get(),
		Player: media.Get(),
		Gate:   func(w *gogui.Window) gogui.View { return app().Gate(w) },
		Redraw: func() { shell.Get().Redraw() },
	})
	apps.Get().Add(dial.App())
	o.Views = append(o.Views, dial.Screen)
}
