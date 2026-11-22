//go:build !board_biscuit

package nws

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/homeassistant/nws"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/remember"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/product"
	featureremember "github.com/franzudev/echolocal/internal/feature/remember"
)

func init() {
	component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(62))
}

var get = sync.OnceValue(func() *nws.NWS {
	return nws.Build(
		product.Get(),
		config.WeatherSection,
		remember.NewSlot[[]string](featureremember.Get(), "weather.dismissed"),
	)
})

func Get() *nws.NWS { return get() }
