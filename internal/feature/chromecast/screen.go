//go:build !board_biscuit

package chromecast

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/display/surface"
	"github.com/franzudev/libcountertop/pkg/display/video"
	"github.com/franzudev/libcountertop/pkg/media/cast"
	"github.com/franzudev/libcountertop/pkg/media/cast/engine"
	castpreferences "github.com/franzudev/libcountertop/pkg/media/cast/preferences"
	"github.com/franzudev/libcountertop/pkg/media/cast/protocols/webreceiver"
	"github.com/franzudev/libcountertop/pkg/media/castreceiver"
	"github.com/franzudev/libcountertop/pkg/media/playback"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/control/castcmd"

	"github.com/franzudev/echolocal/internal/android/dns"
	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/dhcp"
	"github.com/franzudev/echolocal/internal/feature/message"
	"github.com/franzudev/echolocal/internal/hardware/display"
	"github.com/franzudev/echolocal/internal/layout"
	"github.com/franzudev/echolocal/internal/ui"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Engine,
		append([]sharedcomponent.Option{sharedcomponent.Order(59)}, needs()...)...,
	)
}

func configureSettings(*castreceiver.Receiver) {}

func preferences() castpreferences.Provider {
	return castreceiver.Preferences("echolocal", config.CastSection)
}

func needs() []sharedcomponent.Option {
	return []sharedcomponent.Option{component.Needs(board.Panel)}
}

func helper() *surface.Client { return display.Get().Helper() }

func notice() func(title, body string) {
	return func(title, body string) { message.Get().Quiet(title, body, castreceiver.NoticeFor) }
}

var tuning = sync.OnceValue(func() *engine.Tuning {
	t := engine.NewTuning(config.DiagSection)
	component.Settings.Add(t.Controls())
	return t
})

func Tuning() *engine.Tuning { return tuning() }

var host = sync.OnceValue(func() *engine.Host {
	return engine.NewHost(engine.HostOptions{
		Dir:         layout.CastEngineDir,
		State:       layout.CastEngineState,
		Args:        webreceiver.EngineArgs(),
		Nameservers: nameservers,
		Network:     network,
		Screen: func() (int, int) {
			t := pictures()()
			return t.Width, t.Height
		},
		Tuning: Tuning(),
	})
})

func Engine() *engine.Host { return host() }

var page castPage

type castPage struct {
	mu sync.Mutex
	r  *webreceiver.Receiver
}

func (p *castPage) keep(r *webreceiver.Receiver) {
	p.mu.Lock()
	p.r = r
	p.mu.Unlock()
}

func Page() castcmd.Page {
	page.mu.Lock()
	defer page.mu.Unlock()
	if page.r == nil {
		return nil
	}
	return page.r
}

func apps() []cast.Maker {
	return []cast.Maker{
		webreceiver.Maker(webreceiver.Options{
			Engine: Engine(),
			Helper: helper,
			Built:  page.keep,
		}),
	}
}

func nameservers() []string {
	if l := dhcp.Get().Lease(); l != nil && len(l.DNS) > 0 {
		return l.Nameservers()
	}
	return dns.Nameservers()
}

func pictures() func() playback.Target {
	return func() playback.Target {
		return video.Target(display.Get(), component.Board().MaxFPS)
	}
}

func logo() func() []byte { return ui.Logo().PNG }
