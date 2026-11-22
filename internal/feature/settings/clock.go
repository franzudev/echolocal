//go:build !board_biscuit

package settings

import (
	sharedpages "github.com/franzudev/libcountertop/pkg/display/settings"

	"github.com/franzudev/echolocal/internal/feature/dashboard"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

var clocks = sharedpages.NewClockPages(
	sharedpages.ClockDependencies{
		Preferences: preferences(),
		Shell:       shell.Get(),
		Controller:  dashboard.Get(),
		Previews:    previewKit(),
	},
)
