package idle

import (
	"github.com/franzudev/libcountertop/pkg/display/idleview"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/photos"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/visuals"
	"github.com/franzudev/echolocal/internal/hardware/display"
	"github.com/franzudev/echolocal/internal/hardware/gpu"
)

func settings(c config.Config) idleview.Settings {
	return idleview.Settings{Screen: c.Screen, Clock: c.Clock, Idle: c.Idle, Visual: c.Visual}
}

func newView() *idleview.View {
	return idleview.New(idleview.Dependencies{
		Read:    func() idleview.Settings { return settings(config.Get()) },
		Visuals: visuals.Get(), GPU: gpu.Get(), Display: display.Get(), Shell: shell.Get(),
		Photos: photos.Get(),
	})
}
