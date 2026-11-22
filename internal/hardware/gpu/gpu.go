package gpu

import (
	sharedgpu "github.com/franzudev/libcountertop/pkg/display/gpu"

	"github.com/franzudev/echolocal/internal/hardware/display"
)

var renderer = sharedgpu.New(display.Get())

func Get() *sharedgpu.Renderer { return renderer }
