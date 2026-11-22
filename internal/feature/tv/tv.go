//go:build !board_biscuit

package tv

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/homeassistant/yttv"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/chromecast"
	"github.com/franzudev/echolocal/internal/feature/product"
)

func init() { component.Register(sharedcomponent.Network, Get) }

var get = sync.OnceValue(func() *yttv.Guide {
	return yttv.Build(product.Get(), config.TVSection, chromecast.ID)
})

func Get() *yttv.Guide { return get() }
