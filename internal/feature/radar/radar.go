//go:build !board_biscuit

package radar

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/radar"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/product"
)

var get = sync.OnceValue(
	func() *radar.Radar { return radar.Build(product.Get(), config.WeatherSection) },
)

func Get() *radar.Radar { return get() }
