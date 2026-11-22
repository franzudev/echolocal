//go:build !board_biscuit

package gui

import (
	"github.com/franzudev/libcountertop/pkg/display/mirror"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/widgets"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/commands"
	"github.com/franzudev/echolocal/internal/feature/settings"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func extendMirror(look *widgets.UI, o *screens.Options, _ func() *screens.App) {
	if component.Board().CameraWidth == 0 {
		return
	}
	m := mirror.New(mirror.Options{UI: look, Shell: shell.Get(), Camera: settings.Preview()})
	apps.Get().Add(m.App())
	commands.Get().Listen(m.Command())
	o.Views = append(o.Views, m.Screen)
}
