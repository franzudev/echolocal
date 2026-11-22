package network

import (
	"sync"

	"github.com/franzudev/libcountertop/pkg/network/policy"
	sharedwifi "github.com/franzudev/libcountertop/pkg/network/wifi"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/dhcp"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	hwwifi "github.com/franzudev/echolocal/internal/hardware/wifi"
)

func init() {
	component.Register(sharedcomponent.Network, Get, sharedcomponent.Order(5))
}

type Network struct {
	*policy.Policy
	wifi *sharedwifi.Mover
}

var get = sync.OnceValue(func() *Network {
	n := &Network{
		Policy: policy.New(config.NetworkSection),
		wifi:   sharedwifi.NewMover(hwwifi.Get(), dhcp.Get().Fresh),
	}
	n.DownloadThrough(homeassistant.Get().Downloads())
	component.Settings.Add(n.Settings())
	return n
})

func Get() *Network                          { return get() }
func (n *Network) Wifi() *sharedwifi.Mover   { return n.wifi }
func (n *Network) Restore(cfg config.Config) { n.Policy.Restore(cfg.Network.Verify) }
