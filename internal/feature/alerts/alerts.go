package alerts

import (
	"sync"

	shared "github.com/franzudev/libcountertop/pkg/audio/alerts"
	"github.com/franzudev/libcountertop/pkg/audio/tone"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/tones"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
)

func init() { component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(30)) }

var get = sync.OnceValue(func() *shared.Alerts {
	a := shared.New(
		config.AlertsSection,
		component.DevicePlayback,
		tones.Get(),
		func(s tone.Sound, level float64) {
			speaker.Sound().Sound(config.StreamAlerts, level, s)
		},
	)
	configure(a)
	component.Settings.Add(a.Controls)
	tones.Get().Changed.Listen(func(struct{}) { a.Refresh() })
	return a
})

func Get() *shared.Alerts { return get() }
