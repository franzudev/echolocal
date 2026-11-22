//go:build !board_biscuit

package radio

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/homeassistant/radio"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/product"
	"github.com/franzudev/echolocal/internal/layout"
)

var get = sync.OnceValue(func() *radio.Radio {
	return radio.Build(product.Get(), config.RadioSection, func() string {
		mac, _ := layout.FactoryMAC()
		return mac
	})
})

func Get() *radio.Radio { return get() }
