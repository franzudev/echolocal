//go:build !board_biscuit

package gui

import (
	gogui "github.com/go-gui-org/go-gui/gui"

	"github.com/franzudev/libcountertop/pkg/display/agenda"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/calendar"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func extendCalendar(look *widgets.UI, o *screens.Options, app func() *screens.App) {
	cal := calendar.Get().Calendar
	diary := agenda.New(agenda.Options{
		UI:       look,
		Shell:    shell.Get(),
		Calendar: cal,
		Gate:     func(w *gogui.Window) gogui.View { return app().Gate(w) },
		TwentyFour: func() bool {
			return config.Get().Screen.Hours == schema.TwentyFourHour
		},
		Settings: func() *sharedshell.Page { return app().CalendarPage() },
	})
	apps.Get().Add(diary.App())
	o.Calendar = cal
	o.Views = append(o.Views, diary.Screen)
}
