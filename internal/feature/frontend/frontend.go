package frontend

import (
	"context"
	"net/http"
	"sync"

	"github.com/franzudev/libcountertop/pkg/homeassistant/people"
	"github.com/franzudev/libcountertop/pkg/input/textentry"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/webapi"
	"github.com/franzudev/libcountertop/pkg/runtime/webui"

	"github.com/franzudev/echolocal/internal/brand"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/alarm"
	"github.com/franzudev/echolocal/internal/feature/alerts"
	"github.com/franzudev/echolocal/internal/feature/api"
	"github.com/franzudev/echolocal/internal/feature/broadcast"
	"github.com/franzudev/echolocal/internal/feature/call"
	"github.com/franzudev/echolocal/internal/feature/dhcp"
	"github.com/franzudev/echolocal/internal/feature/discovery"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/network"
	"github.com/franzudev/echolocal/internal/feature/voicematch"
	"github.com/franzudev/echolocal/internal/feature/web"
	"github.com/franzudev/echolocal/internal/layout"
)

type componentFrontend struct{ *webapi.API }

func (c componentFrontend) Restore(cfg config.Config) { c.API.Restore(cfg.Frontend) }

var get = sync.OnceValue(func() *webapi.API {
	pages := map[string]http.Handler{
		"alarms":    alarm.Get().Handler(alerts.Get().Tones),
		"call":      call.Get().Handler(discovery.Get().Peers),
		"broadcast": broadcast.Get().WebHandler(),
		"voices":    voicematch.Get().Handler(householdPeople),
	}
	input := &textentry.Controller{}
	addDesktop(pages, input)
	a := webapi.New(webapi.Options{
		Settings: config.FrontendSection,
		Registry: &component.Settings,
		States:   api.Get(),
		Device:   device,
		DeviceID: component.DeviceFrontend,
		Logo:     webapi.Logo{Light: brand.LogoLight, Dark: brand.LogoDark},
		Pages:    pages,
		LogTag:   layout.LogTag,
		Screen:   shot,
		Input:    input,
		Fired:    &component.Fire,
		Address: func() string {
			if lease := dhcp.Get().Lease(); lease != nil {
				return lease.Address.IP.String()
			}
			return ""
		},
		Wifi: network.Get().Wifi(),
	})
	component.Settings.Add(a.Controls())
	return a
})

func Get() *webapi.API { return get() }

var handler = sync.OnceValue(func() http.Handler { return Get().Handler() })

func init() {
	component.Register(
		sharedcomponent.Network,
		func() componentFrontend { return componentFrontend{Get()} },
	)
	web.Get().HandlePublic(
		webapi.Prefix+"/*",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler().ServeHTTP(w, r)
		}),
	)
	web.Get().PublicRoot(webui.Handler())
}

func householdPeople(ctx context.Context) ([]people.Person, error) {
	return people.List(ctx, homeassistant.Get())
}

func device() webapi.Device {
	b := component.Board()
	return webapi.Device{
		ID:      discovery.ID(),
		Name:    config.Get().Device.Name,
		Model:   b.Model,
		Board:   b.Codename,
		Version: layout.Version,
	}
}
