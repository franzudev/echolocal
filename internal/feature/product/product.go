package product

import (
	"sync"

	sharedproduct "github.com/franzudev/libcountertop/pkg/runtime/product"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/states"
	"github.com/franzudev/echolocal/internal/layout"
)

var get = sync.OnceValue(func() sharedproduct.Context {
	return sharedproduct.Context{
		Name:       layout.Manufacturer,
		Version:    layout.Version,
		Site:       "github.com/franzudev/echolocal",
		StateDir:   layout.StateDir,
		Controller: homeassistant.Get(),
		States:     states.Get(),
		Settings:   &component.Settings,
		Redraw:     func() { shell.Get().Redraw() },
		Screen:     component.DeviceScreen,
		Camera:     component.DeviceCamera,
	}
})

func Get() sharedproduct.Context { return get() }
