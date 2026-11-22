//go:build !board_biscuit

package drm

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/display/surface"
	"github.com/franzudev/libcountertop/pkg/media/widevine"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/dhcp"
	"github.com/franzudev/echolocal/internal/hardware/display"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(61),
		component.Needs(board.Panel),
	)
}

var Get = sync.OnceValue(func() *widevine.Provisioner {
	return widevine.New(
		func() *surface.Client { return display.Get().Helper() },
		dhcp.Get(),
		widevine.L1,
	)
})
