package idle

import (
	"sync"
	"time"

	"github.com/franzudev/libcountertop/pkg/display/idleview"
	"github.com/franzudev/libcountertop/pkg/display/panel"
	"github.com/franzudev/libcountertop/pkg/display/videopage"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/brightness"
	"github.com/franzudev/echolocal/internal/feature/drawer"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/timesync"
	"github.com/franzudev/echolocal/internal/feature/web"
	"github.com/franzudev/echolocal/internal/hardware/display"
	hwtouch "github.com/franzudev/echolocal/internal/hardware/touch"
)

type Idle struct{ *idleview.Idle }

var (
	once   sync.Once
	shared *Idle
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(35),
		component.Needs(board.Panel),
	)
}

func Get() *Idle {
	once.Do(func() {
		shared = &Idle{idleview.NewIdle(idleview.Options{
			Settings: config.IdleSection,
			DeviceID: component.DeviceScreen,
			Shell:    shell.Get(),
			View:     newView(),
			Ready:    func() bool { return timesync.Get().Valid() },
			Activity: []idleview.Activity{
				func() (time.Duration, bool) { return hwtouch.Get().Since(), true },
				func() (time.Duration, bool) {
					w := videopage.Watched()
					return time.Since(w), !w.IsZero()
				},
			},
			Free: func() bool {
				return display.Get().Uncovered(panel.PriorityDashboard) &&
					web.Get().Offering() == ""
			},
			Ambient: brightness.Get().Ambient,
		})}
		shared.Night.Listen(brightness.Get().Night)
		component.Settings.Add(shared.Settings())
		drawer.Get().Add(shared.Rail())
	})
	return shared
}

func (i *Idle) Restore(cfg config.Config) { i.Idle.Restore(cfg.Idle) }
