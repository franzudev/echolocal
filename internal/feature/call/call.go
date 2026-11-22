package call

import (
	"sync"

	sharedcall "github.com/franzudev/libcountertop/pkg/media/call"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/detect"
	"github.com/franzudev/echolocal/internal/feature/discovery"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/tones"
	"github.com/franzudev/echolocal/internal/feature/web"
	"github.com/franzudev/echolocal/internal/hardware/buttons"
)

type componentCall struct{ *sharedcall.Calls }

func (c *componentCall) Restore(v config.Config) { c.Calls.Restore(v.Call) }

var once sync.Once
var shared *sharedcall.Calls

func Get() *sharedcall.Calls {
	once.Do(func() {
		d := sharedcall.Dependencies{
			Settings: config.CallSection,
			Peers:    discovery.Get(),
			Port:     web.Port,
			Secure:   web.SecurePort,
			Client:   discovery.Client(),
			Media:    media.Get(),
			Sound:    device{},
			Tones:    tones.Get(),
			Home:     homeassistant.Get(),
		}
		front(&d)
		shared = sharedcall.New(d)
		configureSettings(shared)
		component.Settings.Add(shared.Controls())
	})
	return shared
}

func init() {
	component.Register(
		sharedcomponent.Network,
		func() *componentCall { return &componentCall{Get()} },
		sharedcomponent.Order(62),
	)
	web.HandlePeer("POST /call/offer", discovery.Home(Get().OfferHandler()))
	web.HandlePeer("POST /call/answer", discovery.Home(Get().AnswerHandler()))
	web.HandlePeer("POST /call/end", discovery.Home(Get().EndHandler()))
	// TODO: remove the plain-HTTP call routes once no device predates the secure port.
	web.Handle("POST /call/offer", discovery.Legacy(Get().OfferHandler()))
	web.Handle("POST /call/answer", discovery.Legacy(Get().AnswerHandler()))
	web.Handle("POST /call/end", discovery.Legacy(Get().EndHandler()))
	detect.PauseWhen(Get().PausesWake)
	buttons.Get().Take(take)
}

func take(e buttons.Event) bool {
	if e.Kind != buttons.Tap {
		return false
	}
	c, ok := Get().Now()
	if !ok {
		return false
	}
	switch {
	case c.State == sharedcall.Ringing && e.Name == buttons.Action:
		Get().AnswerAudio()
	case c.State == sharedcall.Ringing && e.Name == buttons.VolumeDown:
		Get().Hangup()
	case (c.State == sharedcall.Talking || c.State == sharedcall.Calling) && e.Name == buttons.Action:
		Get().Hangup()
	default:
		return false
	}
	return true
}
