//go:build !board_biscuit

package weather

import (
	"sync"

	haweather "github.com/ygelfand/libcountertop/pkg/homeassistant/weather"
	sharedcomponent "github.com/ygelfand/libcountertop/pkg/runtime/component"

	"github.com/ygelfand/echolocal/internal/board"
	"github.com/ygelfand/echolocal/internal/component"
	"github.com/ygelfand/echolocal/internal/config"
	openmeteo "github.com/ygelfand/echolocal/internal/feature/openmeteo"
	"github.com/ygelfand/echolocal/internal/feature/product"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(61),
		component.Needs(board.Panel),
	)
}

type Weather struct{ *haweather.Weather }

// Now prefers the Open-Meteo forecast when it is enabled, located and fresh;
// the embedded Home Assistant entity forecast answers otherwise.
func (w Weather) Now() (haweather.Reading, bool) {
	if r, ok := openmeteo.Get().Now(); ok {
		return r, true
	}
	return w.Weather.Now()
}

func (w Weather) Restore(cfg config.Config) { w.Publish(cfg.Weather) }

var get = sync.OnceValue(func() *Weather {
	return &Weather{haweather.Build(product.Get(), config.WeatherSection, config.HomeSection)}
})

func Get() *Weather { return get() }
