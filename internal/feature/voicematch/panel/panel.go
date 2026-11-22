//go:build !board_biscuit

package panel

import (
	"sync"
	"time"

	sharedmatch "github.com/franzudev/libcountertop/pkg/assistant/voicematch"
	sharedscreen "github.com/franzudev/libcountertop/pkg/display/assistant"
	"github.com/franzudev/libcountertop/pkg/display/avatar"
	"github.com/franzudev/libcountertop/pkg/display/voicematchview"
	"github.com/franzudev/libcountertop/pkg/display/widgets"
	"github.com/franzudev/libcountertop/pkg/fetch"
	ha "github.com/franzudev/libcountertop/pkg/homeassistant"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/assistant"
	"github.com/franzudev/echolocal/internal/feature/cameras"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/visuals"
	"github.com/franzudev/echolocal/internal/feature/voicematch"
	"github.com/franzudev/echolocal/internal/hardware/display"
	"github.com/franzudev/echolocal/internal/hardware/gpu"
)

const pictureWait = 15 * time.Second

var People = sync.OnceValue(func() *avatar.Book {
	b := avatar.NewBook(avatar.HomeAssistant{
		Entities: homeassistant.Get(),
		Site:     cameras.Site(),
		Client:   fetch.Client(pictureWait),
	}, shell.Get().Redraw)
	homeassistant.Get().Changed.Listen(func(a ha.Access) {
		if a == ha.Allowed {
			b.Refresh()
		}
	})
	return b
})

var Get = sync.OnceValue(func() *voicematchview.View {
	return voicematchview.New(voicematchview.Options{
		UI:      widgets.NewUI(config.ScreenSection),
		Shell:   shell.Get(),
		Voices:  voicematch.Get(),
		People:  People(),
		Changed: &voicematch.Get().Changed,
		GPU:     gpu.Get(),
		Display: display.Get(),
		Visuals: visuals.Get(),
	})
})

var Follow = sync.OnceFunc(func() {
	voicematch.Get().Identified.Listen(func(i sharedmatch.Identified) {
		if !i.Known {
			return
		}
		People().People()
		p := i.Match.Profile
		assistant.Get().Spoke(sharedscreen.Speaker{ID: p.Who(), Name: p.Name}, i.At)
	})
})
