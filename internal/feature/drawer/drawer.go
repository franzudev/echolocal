package drawer

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/display/drawer"
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/screen"
	"github.com/franzudev/echolocal/internal/feature/shell"
	hwtouch "github.com/franzudev/echolocal/internal/hardware/touch"
)

func init() {
	component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(35))
}

var get = sync.OnceValue(func() *drawer.Rail {
	r := drawer.New(drawer.Options{
		Shell: shell.Get(),
		Edge:  func() schema.Edge { return config.Get().Screen.Drawer },
		Hide:  func() schema.Delay { return config.Get().Screen.DockHide },
	})
	hwtouch.Get().Gestures.Listen(r.Gesture)
	shell.Get().Changed.Listen(func(sharedshell.Change) { r.Settle() })
	screen.Get().Restyled.Listen(func(schema.Screen) { r.Settle() })
	r.Settle()
	return r
})

func Get() *drawer.Rail { return get() }
