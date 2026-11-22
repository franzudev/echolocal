//go:build !board_biscuit

package settings

import (
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"

	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/ui/visual"
)

func lookPages() *sharedsettings.LookPages {
	return sharedsettings.NewLookPages(
		sharedsettings.LookDependencies{
			Preferences: preferences(),
			Shell:       shell.Get(),
			Thumbnail:   visual.Thumbs().Fit,
			Previews:    previewKit(),
		},
	)
}
