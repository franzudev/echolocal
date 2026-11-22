//go:build !board_biscuit

package settings

import (
	sharedpages "github.com/franzudev/libcountertop/pkg/display/settings"
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"

	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/idle"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/visuals"
	"github.com/franzudev/echolocal/internal/feature/weather"
	"github.com/franzudev/echolocal/internal/ui/visual"
)

func idlePages() *sharedpages.IdlePages {
	return sharedpages.NewIdlePages(
		sharedpages.IdleDependencies{
			Preferences:  preferences(),
			Home:         homeassistant.Get(),
			Marks:        modeMarks(),
			Shell:        shell.Get(),
			Idle:         idle.Get(),
			Visuals:      visuals.Get(),
			Weather:      weather.Get(),
			SaverWeather: weather.Get().Saver(),
			Thumbnail:    visual.Thumbs().Fit,
			Previews:     previewKit(),
		},
	)
}

func WeatherEntities() *sharedshell.Page { return idlePages().WeatherEntities() }
