//go:build !board_biscuit

package gui

import (
	"github.com/franzudev/libcountertop/pkg/audio/resample"
	"github.com/franzudev/libcountertop/pkg/display/musicapp"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/video"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/display/youtubeapp"
	"github.com/franzudev/libcountertop/pkg/fetch"
	"github.com/franzudev/libcountertop/pkg/media/playback"
	"github.com/franzudev/libcountertop/pkg/media/youtube"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/apps"
	"github.com/franzudev/echolocal/internal/feature/playout"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/hardware/display"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
)

func extendYouTube(look *widgets.UI, o *screens.Options, _ func() *screens.App) {
	client := fetch.Client(musicapp.FetchTime)
	r := youtube.NewResolver(client, func() playback.Target {
		return video.Target(display.Get(), component.Board().MaxFPS)
	})
	r.Resample = func(from int) func([]int16) []int16 {
		return resample.NewRational(from, speaker.Rate, speaker.Channels).Run
	}
	yt := youtubeapp.New(youtubeapp.Options{
		UI:       look,
		Shell:    shell.Get(),
		Accounts: config.YouTubeAccountsSection,
		Resolver: r,
		Output:   playout.New(),
		Skip:     func() []string { return config.Get().Cast.YouTube.Skip },
		Client:   client,
		Redraw:   func() { shell.Get().Redraw() },
	})
	apps.Get().Add(yt.App())
	o.Views = append(o.Views, yt.Screen)
}
