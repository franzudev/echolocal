package remember

import (
	"log/slog"
	"sync"

	"github.com/franzudev/libcountertop/pkg/runtime/remember"

	"github.com/franzudev/echolocal/internal/layout"
)

var get = sync.OnceValue(func() *remember.Store {
	s, err := remember.Open(layout.RuntimePath)
	if err != nil {
		slog.Warn("reading runtime state", "err", err)
	}
	return s
})

func Get() *remember.Store { return get() }
