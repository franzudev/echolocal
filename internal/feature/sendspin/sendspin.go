package sendspin

import (
	"log/slog"
	"sync"

	sharedplayer "github.com/franzudev/libcountertop/pkg/media/sendspin"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/android/firewall"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/timesync"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
)

type componentPlayer struct{ *sharedplayer.Player }

func (p *componentPlayer) Restore(c config.Config) { p.Player.Restore(c.Sendspin) }

var get = sync.OnceValue(func() *sharedplayer.Player {
	p := sharedplayer.New(sharedplayer.Dependencies{
		Output:      newOutput(),
		Arbitration: speaker.Sound().Backgrounds(),
		Device:      identity{},
		Settings:    config.SendspinSection,
		DeviceID:    component.DevicePlayback,
		Clock:       timesync.Get(),
		Media:       media.Get(),
		Queue:       queue(),
		Favorite:    favorite(),
		OpenPort: func() {
			if err := firewall.Open(firewall.Sendspin, Port); err != nil {
				slog.Error("opening the sendspin port failed", "port", Port, "err", err)
			}
		},
	})
	component.Settings.Add(p.Settings())
	return p
})

func Get() *sharedplayer.Player { return get() }

func ClientID() string { return identity{}.Identity().ID }

func init() {
	component.Register(
		sharedcomponent.Device,
		func() *componentPlayer { return &componentPlayer{Get()} },
		sharedcomponent.Order(26),
	)
}
