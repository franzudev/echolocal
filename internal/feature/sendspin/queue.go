//go:build !board_biscuit

package sendspin

import (
	"github.com/franzudev/libcountertop/pkg/homeassistant/musicassistant"
	sharedplayer "github.com/franzudev/libcountertop/pkg/media/sendspin"

	"github.com/franzudev/echolocal/internal/feature/product"
)

func queue() sharedplayer.Queue {
	return musicassistant.Build(product.Get(), ClientID)
}

func favorite() sharedplayer.Favorite {
	return musicassistant.BuildLikes(product.Get(), ClientID)
}
