package voicematch

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/assistant/voicematch"
	ha "github.com/franzudev/libcountertop/pkg/homeassistant"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/hardware/mic"
	"github.com/franzudev/echolocal/internal/layout"
)

func init() { component.Register(sharedcomponent.Device, Get) }

const EnrolmentEvent = "esphome.echolocal_enrolment"

var get = sync.OnceValue(func() *voicematch.Service {
	s := voicematch.NewService(voicematch.Options{
		Speech:   mic.Get().History,
		ModelDir: layout.VoiceprintDir,
		Profiles: layout.VoicesPath,
		Outbox:   layout.VoicesOutbox,
		Home:     homeassistant.Get(),
		Emit: func(fields map[string]string) {
			component.Fire.Emit(sharedcomponent.Event{Name: EnrolmentEvent, Data: fields})
		},
	})
	homeassistant.Get().Changed.Listen(func(a ha.Access) {
		if a == ha.Allowed {
			s.Resync()
		}
	})
	return s
})

func Get() *voicematch.Service { return get() }
