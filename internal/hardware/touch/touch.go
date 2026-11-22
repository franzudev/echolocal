package touch

import (
	"sync"
	"time"

	"github.com/franzudev/libcountertop/pkg/input/touch"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/service"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/hardware/display"
)

func init() {
	component.Register(sharedcomponent.Hardware, Get, sharedcomponent.Order(30),
		component.Needs(board.Panel),
		sharedcomponent.Supervise(service.Restart(time.Second, time.Minute)))
}

var get = sync.OnceValue(
	func() *touch.Screen { return touch.New(component.Board().Touch, display.Get()) },
)

func Get() *touch.Screen { return get() }
