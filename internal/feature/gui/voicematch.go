//go:build !board_biscuit

package gui

import (
	"github.com/franzudev/libcountertop/pkg/display/screens"
	"github.com/franzudev/libcountertop/pkg/display/widgets"

	"github.com/franzudev/echolocal/internal/feature/voicematch"
	"github.com/franzudev/echolocal/internal/feature/voicematch/panel"
)

func extendVoiceMatch(_ *widgets.UI, o *screens.Options, _ func() *screens.App) {
	panel.Follow()
	o.Views = append(o.Views, panel.Get().Screen)
	o.Redraws = append(o.Redraws, screens.On(&voicematch.Get().Changed))
	o.Badge = panel.People().Badge
}
