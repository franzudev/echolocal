//go:build !board_biscuit

package settings

import (
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"

	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/triggers"
)

func voicePage() *sharedshell.Page {
	return sharedsettings.VoicePage(triggers.Get().Controls(), homeassistant.Get(), shell.Get())
}
