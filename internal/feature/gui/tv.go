//go:build !board_biscuit

package gui

import (
	gogui "github.com/go-gui-org/go-gui/gui"

	"github.com/franzudev/libcountertop/pkg/display/musicapp"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	libshell "github.com/franzudev/libcountertop/pkg/display/shell"
	"github.com/franzudev/libcountertop/pkg/display/tvapp"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/fetch"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/shell"
	featuretv "github.com/franzudev/echolocal/internal/feature/tv"
	"github.com/franzudev/echolocal/internal/feature/volume"
)

func extendTV(look *widgets.UI, o *screens.Options, app func() *screens.App) {
	tv := tvapp.New(tvapp.Options{
		UI:      look,
		Shell:   shell.Get(),
		Guide:   featuretv.Get(),
		Player:  media.Get(),
		Loud:    func() { volume.Get().Show(config.StreamMedia) },
		Client:  fetch.Client(musicapp.FetchTime),
		Gate:    func(w *gogui.Window) gogui.View { return app().Gate(w) },
		Clock24: func() bool { return config.Get().Screen.Hours == schema.TwentyFourHour },
		Redraw:  func() { shell.Get().Redraw() },
	})
	apps.Get().Add(tv.App())
	o.Views = append(o.Views, tv.Screen)
	o.YouTubeTVPage = func() libshell.View { return tv.Page() }
}
