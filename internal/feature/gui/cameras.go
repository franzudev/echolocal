//go:build !board_biscuit

package gui

import (
	"github.com/franzudev/libcountertop/pkg/display/cameraview"
	"github.com/franzudev/libcountertop/pkg/display/doorbellview"
	"github.com/franzudev/libcountertop/pkg/display/homealarmview"
	"github.com/franzudev/libcountertop/pkg/display/photosview"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/shell"
	"github.com/franzudev/libcountertop/pkg/display/thermostatview"
	"github.com/franzudev/libcountertop/pkg/display/widgets"

	"github.com/franzudev/echolocal/internal/hardware/display"
	"github.com/franzudev/echolocal/internal/hardware/gpu"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/cameras"
	"github.com/franzudev/echolocal/internal/feature/doorbell"
	"github.com/franzudev/echolocal/internal/feature/homealarm"
	"github.com/franzudev/echolocal/internal/feature/homecontrol"
	"github.com/franzudev/echolocal/internal/feature/photos"
	"github.com/franzudev/echolocal/internal/feature/settings"
	featureshell "github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/thermostat"
)

func extendCameras(ui *widgets.UI, o *screens.Options, _ func() *screens.App) {
	apps.Get().Add(cameras.Get().App())
	apps.Get().Add(doorbell.View().App())
	photosApp := photosview.New(photosview.Options{
		Shell:    featureshell.Get(),
		Photos:   photos.Get(),
		Settings: config.PhotosSection,
		Page:     func() shell.View { return settings.Visuals().Photos() },
		GPU:      gpu.Get(),
		Display:  display.Get(),
	})
	apps.Get().Add(photosApp.App())
	alarmPanel := homealarm.NewView(ui)
	apps.Get().Add(alarmPanel.App())
	o.HomeAlarmPage = func() shell.View { return alarmPanel.Page() }
	o.Views = append(o.Views, alarmPanel.Screen)
	climate := thermostat.NewView(ui)
	apps.Get().Add(climate.App())
	o.ThermostatPage = func() shell.View { return climate.Page() }
	o.Views = append(o.Views, climate.Screen)
	for _, a := range homecontrol.Get().Apps() {
		apps.Get().Add(a)
	}
	o.Views = append(o.Views, cameras.Get().Screen, doorbell.View().Screen, photosApp.Screen)
	if o.Boards == nil {
		o.Boards = map[string]screens.Board{}
	}
	o.Boards[cameraview.TabKind] = screens.Board{
		Glyph: cameras.Get().Glyph,
		Draw:  cameras.Get().Board,
	}
	o.Boards[doorbellview.TabKind] = screens.Board{
		Glyph: doorbell.View().Glyph,
		Draw:  doorbell.View().Board,
	}
	o.Boards[homealarmview.TabKind] = screens.Board{
		Glyph: alarmPanel.Glyph,
		Draw:  alarmPanel.Board,
	}
	o.Boards[thermostatview.TabKind] = screens.Board{Glyph: climate.Glyph, Draw: climate.Board}
}
