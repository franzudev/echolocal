package message

import (
	"sync"

	shared "github.com/franzudev/libcountertop/pkg/display/message"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/alerts"
)

var get = sync.OnceValue(func() *shared.Messages {
	m := shared.New()
	m.SoundWith(alerts.Get())
	return m
})

func Get() *shared.Messages { return get() }
func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(40),
		component.Needs(board.Panel),
	)
}
