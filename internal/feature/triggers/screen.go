//go:build !board_biscuit

package triggers

import (
	setting "github.com/franzudev/libcountertop/pkg/settings"
	"github.com/franzudev/libcountertop/pkg/voice"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/commands"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
)

func commandNames() []string {
	if !component.Board().Has(board.Panel) {
		return nil
	}
	out := []string{"radar"}
	if component.Board().CameraWidth > 0 {
		out = append(out, "mirror")
	}
	return out
}

func applyCommands() { commands.Get().Apply() }

func labelChoices() func() []setting.Option { return voice.LabelChoices(homeassistant.Get()) }
