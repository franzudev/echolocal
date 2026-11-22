//go:build !board_biscuit

package gui

import (
	"sync"
	"time"

	"github.com/franzudev/echolocal/internal/hardware/gpu"

	sharedvolume "github.com/franzudev/libcountertop/pkg/audio/volume"
	"github.com/franzudev/libcountertop/pkg/display/screens"
	sharedsettings "github.com/franzudev/libcountertop/pkg/display/settings"
	"github.com/franzudev/libcountertop/pkg/display/shell"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/media/lyrics"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/service"
	"github.com/franzudev/libcountertop/pkg/say"
	tesseraview "github.com/franzudev/libcountertop/pkg/tessera/widgets"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/a2dp"
	"github.com/franzudev/echolocal/internal/feature/assistant"
	"github.com/franzudev/echolocal/internal/feature/cameras"
	"github.com/franzudev/echolocal/internal/feature/chromecast"
	"github.com/franzudev/echolocal/internal/feature/dashboard"
	"github.com/franzudev/echolocal/internal/feature/doorbell"
	"github.com/franzudev/echolocal/internal/feature/drawer"
	"github.com/franzudev/echolocal/internal/feature/firmware"
	"github.com/franzudev/echolocal/internal/feature/frontend"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/homecontrol"
	"github.com/franzudev/echolocal/internal/feature/idle"
	"github.com/franzudev/echolocal/internal/feature/lovelace"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/message"
	"github.com/franzudev/echolocal/internal/feature/nws"
	"github.com/franzudev/echolocal/internal/feature/poster"
	"github.com/franzudev/echolocal/internal/feature/privacy"
	"github.com/franzudev/echolocal/internal/feature/product"
	"github.com/franzudev/echolocal/internal/feature/screen"
	"github.com/franzudev/echolocal/internal/feature/settings"
	featureshell "github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/tessera"
	"github.com/franzudev/echolocal/internal/feature/timer"
	"github.com/franzudev/echolocal/internal/feature/volume"
	"github.com/franzudev/echolocal/internal/feature/weather"
	"github.com/franzudev/echolocal/internal/feature/web"
	"github.com/franzudev/echolocal/internal/hardware/display"
	hwtouch "github.com/franzudev/echolocal/internal/hardware/touch"
	"github.com/franzudev/echolocal/internal/lib/wake"
	echoui "github.com/franzudev/echolocal/internal/ui"
	"github.com/franzudev/echolocal/internal/ui/theme"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(36),
		sharedcomponent.Supervise(service.Restart(time.Second, time.Minute)),
		component.Needs(board.Panel),
	)
}

var Get = sync.OnceValue(func() *screens.App {
	look := widgets.NewUI(config.ScreenSection)
	o := screens.Options{
		Name:       "echolocal",
		UISize:     component.Board().UISize,
		DeviceName: func() string { return config.Get().Device.Name },
		Settings: func() screens.Settings {
			c := config.Get()
			return screens.Settings{
				Screen:  c.Screen,
				Clock:   c.Clock,
				Idle:    c.Idle,
				Weather: c.Weather,
				Media:   c.Media,
			}
		},
		UI:           look,
		Weather:      config.WeatherSection,
		Brand:        theme.Brand,
		Logo:         echoui.Logo().Light,
		Night:        echoui.Logo().Dark,
		Offering:     web.Get().Offering,
		Input:        frontend.Get().Input(),
		KeyboardPair: frontend.Get().KeyboardPair,

		Display:   display.Get(),
		GPU:       gpu.Get(),
		Touch:     hwtouch.Get(),
		Shell:     featureshell.Get(),
		Screen:    screen.Get().Screen,
		Messages:  message.Get(),
		Assistant: assistant.Get(),
		Stage:     assistant.Stage(),
		Timers:    timer.Get(),
		Poster:    poster.Get(),
		Drawer:    drawer.Get(),
		Tabs:      dashboard.Tabs(),
		Home:      homecontrol.Get(),
		HomeHA:    homeassistant.Get(),
		Adopted:   func() bool { return config.Get().API.Adopted },
		Saver:     idle.Get(),
		Visuals:   settings.Visuals(),
		Tessera: tesseraview.New(tesseraview.Dependencies{
			UI:      look,
			Screens: tessera.Get(),
			Perform: homeassistant.Get().Perform,

			HomeAfterIdle: homecontrol.Get().ReturnHome,
			Logo:          echoui.Logo(),
		}),
		Firmware: tessera.Firmware(),
		Lovelace: lovelace.Get(),
		Forecast: weather.Get(),
		Warnings: nws.Get(),

		WeatherEntities: func() shell.View { return settings.WeatherEntities() },
		CamerasPage:     func() shell.View { return cameras.Get().Page() },
		DoorbellPage:    func() shell.View { return doorbell.View().Page() },
		Player:          media.Get(),
		PlayerPage:      media.Page,
		BluetoothSources: func() shell.View {
			return sharedsettings.BluetoothSourcesPage(a2dp.Get(), featureshell.Get())
		},
		Lyrics:      lyrics.Build(product.Get()),
		MediaVolume: sharedvolume.For(volume.Get(), config.StreamMedia),
		Volume:      volume.Get(),
		Streams:     config.Streams(),

		Upgrade: func() widgets.Upgrade {
			if up := firmware.Get().Upgrade(); up.Active() {
				return widgets.Upgrade{Active: true, At: up.At, Version: up.Version}
			}
			if in := chromecast.Engine().Installation(); in.Active {
				return widgets.Upgrade{Active: true, At: in.At, Label: say.T("upgrade.engine")}
			}
			p := wake.Lib().Preparation()
			return widgets.Upgrade{Active: p.Active(), At: p.At, Label: p.Phrase}
		},
		Upgrading: screens.All(
			screens.On(&firmware.Get().Upgrading),
			screens.On(&chromecast.Engine().Installing),
			screens.On(&wake.Lib().Preparing),
		),
		Marks: func() screens.Marks {
			m := privacy.Get().Current()
			return screens.Marks{Mic: m.MicMuted, Camera: m.CameraBlocked}
		},
		Redraws: []screens.Watch{
			screens.On(&firmware.Get().Offered),
			screens.On(&firmware.Get().Listed),
			screens.On(&privacy.Get().Changed),
			screens.On(&web.Get().Offered),
		},
	}
	var app *screens.App
	self := func() *screens.App { return app }
	for _, extend := range []func(*widgets.UI, *screens.Options, func() *screens.App){
		extendAlarms,
		extendApps,
		extendCalendar,
		extendCalls,
		extendCameras,
		extendLock,
		extendMirror,
		extendMusic,
		extendRadar,
		extendRadio,
		extendShopping,
		extendSync,
		extendTV,
		extendVoiceMatch,
		extendYouTube,
	} {
		extend(look, &o, self)
	}
	app = screens.New(o)
	return app
})
