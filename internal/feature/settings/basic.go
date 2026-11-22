//go:build !board_biscuit

package settings

import (
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"
	"github.com/franzudev/libcountertop/pkg/display/syncview"

	"github.com/franzudev/echolocal/internal/feature/access"
	"github.com/franzudev/echolocal/internal/feature/alerts"
	"github.com/franzudev/echolocal/internal/feature/chromecast"
	"github.com/franzudev/echolocal/internal/feature/dashboard"
	"github.com/franzudev/echolocal/internal/feature/firmware"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/homecontrol"
	"github.com/franzudev/echolocal/internal/feature/lock"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/network"
	"github.com/franzudev/echolocal/internal/feature/photos"
	"github.com/franzudev/echolocal/internal/feature/poster"
	"github.com/franzudev/echolocal/internal/feature/privacy"
	"github.com/franzudev/echolocal/internal/feature/screen"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/timesync"
	"github.com/franzudev/echolocal/internal/feature/visuals"
)

func Visuals() *sharedsettings.BasicPages { return basicPages() }

func basicPages() *sharedsettings.BasicPages {
	return sharedsettings.NewBasicPages(
		sharedsettings.BasicDependencies{
			Preferences: preferences(),
			Shell:       shell.Get(),
			Screen:      screen.Get(),
			Poster:      poster.Get(),
			Photos:      photos.Get(),
			Visuals:     visuals.Get(),
			Engine:      chromecast.Tuning(),
			Privacy:     privacy.Get(),
			Access:      access.Get(),
			Zone:        timesync.Get(),
			Dashboard:   dashboard.Get(),
			Clocks:      clocks,
			Idle:        idlePages(),
			UISize:      uiSize,
			Home:        homecontrol.Get(),
			HA:          homeassistant.Get(),
			Marks:       modeMarks(),
			Alerts:      alerts.Get(),
			Updates:     firmware.Get(),
			Lock:        lock.Get().Gate(),
			HomeKey:     network.Get(),
			Media:       media.Get(),
			Player:      media.Card(),
			Sponsor:     servicePages(),
			Previews:    previewKit(),
			PowerPage:   func() sharedshell.View { return powerPage() },
			SyncPage:    func() sharedshell.View { return &syncview.Directory{} },
			VisualPage:  func() sharedshell.View { return visualPage() },
		},
	)
}
