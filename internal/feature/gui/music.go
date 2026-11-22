//go:build !board_biscuit

package gui

import (
	gogui "github.com/go-gui-org/go-gui/gui"

	"github.com/franzudev/libcountertop/pkg/display/musicapp"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/fetch"
	"github.com/franzudev/libcountertop/pkg/homeassistant/musicassistant"

	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/cameras"
	"github.com/franzudev/echolocal/internal/feature/discovery"
	"github.com/franzudev/echolocal/internal/feature/product"
	"github.com/franzudev/echolocal/internal/feature/sendspin"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/ui"
)

func extendMusic(look *widgets.UI, o *screens.Options, app func() *screens.App) {
	music := musicapp.New(musicapp.Options{
		UI:     look,
		Shell:  shell.Get(),
		Music:  musicassistant.BuildMusic(product.Get(), sendspin.ClientID),
		Peers:  discovery.Get(),
		Logo:   sharedsettings.Themed(ui.Logo().Light, ui.Logo().Dark),
		Site:   cameras.Site().URL,
		Client: fetch.Client(musicapp.FetchTime),
		Gate:   func(w *gogui.Window) gogui.View { return app().Gate(w) },
		Redraw: func() { shell.Get().Redraw() },
	})
	apps.Get().Add(music.App())
	o.Views = append(o.Views, music.Screen)
}
