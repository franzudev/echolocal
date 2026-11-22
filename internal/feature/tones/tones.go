package tones

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/audio/tones"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
	"github.com/franzudev/echolocal/internal/layout"
)

func init() { component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(15)) }

var get = sync.OnceValue(func() *tones.Store {
	s := tones.New(tones.Options{
		Dir:      layout.ToneDir,
		Rate:     speaker.Rate,
		Channels: speaker.Channels,
	})
	s.Changed.Listen(func(struct{}) { component.Reconnect.Emit(struct{}{}) })
	return s
})

func Get() *tones.Store { return get() }
