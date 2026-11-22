package screen

import (
	"sync"

	libscreen "github.com/franzudev/libcountertop/pkg/display/screen"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/hardware/display"
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(10),
		component.Needs(board.Panel),
	)
}

type Screen struct{ *libscreen.Screen }

var get = sync.OnceValue(func() *Screen {
	s := &Screen{libscreen.New(libscreen.Options{
		Settings:  config.ScreenSection,
		DeviceID:  component.DeviceScreen,
		Backlight: func(level int) error { return display.Get().Brightness(level) },
		UISize:    func() string { return component.Board().UISize },
	})}
	component.Settings.Add(s.Settings())
	return s
})

func Get() *Screen { return get() }

func Table() *libscreen.Table { return Get().Table() }

func (s *Screen) Restore(c config.Config) { s.Screen.Restore(c.Screen) }
