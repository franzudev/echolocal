//go:build !board_biscuit

package settings

import (
	sharedcast "github.com/franzudev/libcountertop/pkg/display/settings/castpages"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/chromecast"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func servicePages() *sharedcast.Pages {
	return sharedcast.New(
		sharedcast.Dependencies{
			Settings: config.CastSection,
			Shell:    shell.Get(),
			Receiver: chromecast.Get(),
		},
	)
}
