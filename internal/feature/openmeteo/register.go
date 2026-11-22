//go:build !board_biscuit

package openmeteo

import (
	"context"
	"sync"
	"time"

	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/service"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/shell"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(62),
		sharedcomponent.Supervise(service.Restart(time.Second, time.Minute)),
	)
}

var (
	once   sync.Once
	shared *Source
)

// Get builds the single Open-Meteo source.
func Get() *Source {
	once.Do(func() {
		shared = New(Options{
			Enabled: func() bool { return config.Get().OpenMeteo.Enabled },
			LatText: func() string { return config.Get().OpenMeteo.Latitude },
			LonText: func() string { return config.Get().OpenMeteo.Longitude },
			Agent:   "echolocal",
			Redraw:  func() { shell.Get().Redraw() },
		})
	})
	return shared
}

// Run keeps the poller alive when no one has asked for the source yet.
func Run(ctx context.Context) error { return Get().Run(ctx) }
