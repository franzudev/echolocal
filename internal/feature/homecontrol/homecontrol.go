package homecontrol

import (
	"sync"

	shareddashboard "github.com/franzudev/libcountertop/pkg/display/dashboard"
	sharedhome "github.com/franzudev/libcountertop/pkg/homeassistant/homecontrol"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/dashboard"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/states"
)

var get = sync.OnceValue(func() *sharedhome.Engine {
	e := sharedhome.New(
		sharedhome.Dependencies{
			Settings: config.HomeSection,
			Client:   homeassistant.Get(),
			States:   states.Get(),
			Shell:    shell.Get(),
			Grid: schema.TesseraGrid{
				Columns: component.Board().Tessera.Columns,
				Rows:    component.Board().Tessera.Rows,
			},
		},
	)
	component.Settings.Add(e.Controls())
	return e
})

func Get() *sharedhome.Engine { return get() }
func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(6),
		component.Needs(board.Panel),
	)
	dashboard.Tabs().Add(10, func() []shareddashboard.Tab { return Get().DashboardTabs() })
}
