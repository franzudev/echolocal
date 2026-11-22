//go:build !board_biscuit

package settings

import (
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/call"
	"github.com/franzudev/echolocal/internal/feature/discovery"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func callsPage() *sharedshell.Page { return callPages().Calls() }

func callPages() *sharedsettings.CallPages {
	return sharedsettings.NewCallPages(
		sharedsettings.CallDependencies{
			Settings:   config.CallSection,
			Shell:      shell.Get(),
			Controller: call.Get(),
			Peers:      discovery.Get().Peers,
			Home:       homeassistant.Get(),
		},
	)
}
