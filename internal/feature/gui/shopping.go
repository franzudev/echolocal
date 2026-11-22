//go:build !board_biscuit

package gui

import (
	gogui "github.com/go-gui-org/go-gui/gui"

	"github.com/franzudev/libcountertop/pkg/display/screens"
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"
	"github.com/franzudev/libcountertop/pkg/display/shopping"
	"github.com/franzudev/libcountertop/pkg/display/widgets"

	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/todo"
)

func extendShopping(look *widgets.UI, o *screens.Options, app func() *screens.App) {
	list := todo.Get().Todo
	shop := shopping.New(shopping.Options{
		UI:       look,
		Shell:    shell.Get(),
		Todo:     list,
		Gate:     func(w *gogui.Window) gogui.View { return app().Gate(w) },
		Settings: func() *sharedshell.Page { return app().TodoLists() },
	})
	apps.Get().Add(shop.App())
	o.Todo = list
	o.Views = append(o.Views, shop.Screen)
}
