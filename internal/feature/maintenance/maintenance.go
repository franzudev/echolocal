package maintenance

import (
	"context"
	"sync"
	"time"

	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/recording"
)

const Period = 5 * time.Minute

func init() {
	component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(80))
}

var (
	once   sync.Once
	shared *Maintenance
)

type Maintenance struct{}

func Get() *Maintenance {
	once.Do(func() { shared = &Maintenance{} })
	return shared
}

func (m *Maintenance) Name() string { return "maintenance" }

func (m *Maintenance) Run(ctx context.Context) error {
	t := time.NewTicker(Period)
	defer t.Stop()

	for {
		recording.Get().Prune()

		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
		}
	}
}
