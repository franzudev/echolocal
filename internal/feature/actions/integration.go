//go:build !board_biscuit

package actions

import (
	sharedha "github.com/franzudev/libcountertop/pkg/homeassistant"

	"github.com/franzudev/echolocal/internal/feature/states"
)

func (a *Actions) followIntegration() {
	states.Get().Follow(sharedha.VersionEntity, "").Reported.Listen(func(string) { a.Declare() })
}
