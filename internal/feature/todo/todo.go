//go:build !board_biscuit

package todo

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/homeassistant/todo"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/product"
)

func init() {
	component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(61))
}

type Todo struct{ *todo.Todo }

func (t Todo) Restore(cfg config.Config) { t.Publish(cfg.Todo) }

var get = sync.OnceValue(
	func() *Todo { return &Todo{todo.Build(product.Get(), config.TodoSection)} },
)

func Get() *Todo { return get() }
