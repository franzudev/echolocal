package vision

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/camera/vision"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/livecam"
	"github.com/franzudev/echolocal/internal/hardware/buttons"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(40),
		component.Needs(board.Camera),
	)
}

var Get = sync.OnceValue(func() *vision.Vision {
	return vision.New(vision.Options{
		DeviceID: component.DeviceCamera,
		Covered:  func() bool { return buttons.Get().Covered() },
		Picture:  livecam.Get().Still,
	})
})
