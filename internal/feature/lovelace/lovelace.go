//go:build !board_biscuit

package lovelace

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/homeassistant/lovelaceview"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/chromecast"
	"github.com/franzudev/echolocal/internal/feature/product"
	"github.com/franzudev/echolocal/internal/layout"
)

func init() {
	component.Register(sharedcomponent.Network, Get, component.Needs(board.Panel))
}

var Get = sync.OnceValue(func() *lovelaceview.Feature {
	return lovelaceview.Build(
		product.Get(),
		config.LovelaceSection,
		chromecast.Engine(),
		func() string {
			return layout.Slug(config.Get().Device.Name)
		},
	)
})
