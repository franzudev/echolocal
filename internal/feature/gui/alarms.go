//go:build !board_biscuit

package gui

import (
	"time"

	sharedalarm "github.com/franzudev/libcountertop/pkg/assistant/alarm"
	"github.com/franzudev/libcountertop/pkg/display/alarms"
	"github.com/franzudev/libcountertop/pkg/display/dawnview"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/alarm"
	"github.com/franzudev/echolocal/internal/feature/alerts"
	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/brightness"
	"github.com/franzudev/echolocal/internal/feature/idle"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/visuals"
)

func extendAlarms(_ *widgets.UI, o *screens.Options, _ func() *screens.App) {
	wake := alarms.New(alarms.Options{
		Shell:  shell.Get(),
		Alarms: alarm.Get(),
		Tones:  alerts.Get().Tones,
	})
	apps.Get().Add(wake.App())
	o.Alarms = alarm.Get()
	light := dawnview.NewLight(
		shell.Get(),
		alarm.Get(),
		visuals.Show(dawnview.NewSource(dawnview.Options{
			Alarms:  alarm.Get(),
			Look:    func() schema.WakeLook { return alerts.Get().Current().Wake },
			Clock24: func() bool { return config.Get().Screen.Hours == schema.TwentyFourHour },
		})),
		idle.Get(),
	)
	alarm.Get().Dawning.Listen(func(sharedalarm.Dawn) { light.Check() })
	alarm.Get().Changed.Listen(func(struct{}) {
		if alarm.Get().Ringing() {
			idle.Get().Wake()
		}
		light.Check()
		shell.Get().Redraw()
	})
	brightness.Get().Rising(func() (float64, bool) {
		if d := alarm.Get().Dawn(); d.Rising() {
			return d.Share(time.Now()), true
		}
		return 1, dawnview.Showing(alarm.Get())
	})
}
