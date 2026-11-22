//go:build !board_biscuit

package thermostat

import (
	"sync"
	"sync/atomic"

	esphome "github.com/ygelfand/go-esphome-device"

	"github.com/franzudev/libcountertop/pkg/display/thermostatview"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	shared "github.com/franzudev/libcountertop/pkg/homeassistant/thermostat"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/dashboard"
	"github.com/franzudev/echolocal/internal/feature/product"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/weather"
)

func init() {
	component.Register(
		sharedcomponent.Network,
		func() *Thermostats { return &Thermostats{Get()} },
		component.Needs(board.Panel),
	)
}

var Get = sync.OnceValue(func() *shared.Thermostats {
	return shared.Build(product.Get(), config.ThermostatSection, func() string {
		r, _ := weather.Get().Now()
		return r.Unit
	})
})

var shown atomic.Pointer[thermostatview.View]

func Shown() *thermostatview.View { return shown.Load() }

func NewView(ui *widgets.UI) *thermostatview.View {
	v := thermostatview.New(thermostatview.Options{
		UI:          ui,
		Shell:       shell.Get(),
		Thermostats: Get(),
		Settings:    config.ThermostatSection,
		Clock24:     func() bool { return config.Get().Screen.Hours == schema.TwentyFourHour },
		Outside: func() (float64, bool) {
			r, ok := weather.Get().Now()
			if !ok || r.Temperature == nil {
				return 0, false
			}
			return *r.Temperature, true
		},
	})
	shown.Store(v)
	dashboard.Tabs().Add(45, v.DashboardTabs)
	return v
}

type Thermostats struct{ *shared.Thermostats }

func (t *Thermostats) Entities() []esphome.Entity { return t.Thermostats.Entities() }

func (t *Thermostats) Restore(cfg config.Config) { t.Thermostats.Restore(cfg.Thermostat) }
