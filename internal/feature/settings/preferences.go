//go:build !board_biscuit

package settings

import (
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"

	"github.com/franzudev/echolocal/internal/config"
)

func preferences() sharedsettings.Preferences {
	return sharedsettings.Preferences{
		Screen:  config.ScreenSection,
		Clock:   config.ClockSection,
		Idle:    config.IdleSection,
		Visual:  config.VisualSection,
		Poster:  config.PosterSection,
		Weather: config.WeatherSection,
		Wake:    config.WakeSection,
		Call:    config.CallSection,
	}
}
