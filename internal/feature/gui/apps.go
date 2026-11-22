//go:build !board_biscuit

package gui

import (
	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/widgets"

	"github.com/franzudev/echolocal/internal/feature/apps"
)

func extendApps(look *widgets.UI, o *screens.Options, _ func() *screens.App) {
	o.Views = append(o.Views, apps.Get().Screens(look))
}
