//go:build !board_biscuit

package gui

import (
	"time"

	gogui "github.com/go-gui-org/go-gui/gui"

	"github.com/franzudev/libcountertop/pkg/display/radarview"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/commands"
	"github.com/franzudev/echolocal/internal/feature/radar"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func extendRadar(look *widgets.UI, o *screens.Options, app func() *screens.App) {
	sky := radarview.New(radarview.Options{
		UI:       look,
		Shell:    shell.Get(),
		Radar:    radar.Get(),
		Settings: config.WeatherSection,
		Clock24: func() bool {
			return config.Get().Screen.Hours == schema.TwentyFourHour
		},
		Again: func(w *gogui.Window, after time.Duration) { app().Again(w, after) },
		Gate:  func(w *gogui.Window) gogui.View { return app().Gate(w) },
	})
	apps.Get().Add(sky.App())
	commands.Get().Listen(sky.Command())
	o.OpenRadar = sky.Open
	o.Views = append(o.Views, sky.Screen)
}
