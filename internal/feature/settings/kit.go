//go:build !board_biscuit

package settings

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/display/previews"
	"github.com/franzudev/libcountertop/pkg/display/theme"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	schema "github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/config"
)

var previewKit = sync.OnceValue(func() previews.Kit {
	ui := widgets.NewUI(config.ScreenSection)
	return previews.NewKit(ui,
		func() bool { return config.Get().Screen.Hours == schema.TwentyFourHour },
		func() theme.Theme { return config.Get().Clock.Ink.Over(ui.Palette()) },
	)
})
