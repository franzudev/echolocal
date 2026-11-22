//go:build !board_biscuit

package settings

import (
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"
	tesseraview "github.com/franzudev/libcountertop/pkg/tessera/widgets"

	echoui "github.com/franzudev/echolocal/internal/ui"
)

func modeMarks() sharedsettings.ModeMarks {
	return sharedsettings.ModeMarks{
		Tessera: tesseraview.Mark,
		Native:  sharedsettings.Themed(echoui.Logo().Light, echoui.Logo().Dark),
	}
}
