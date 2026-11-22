package homeassistant

import (
	"sync"

	sharedha "github.com/franzudev/libcountertop/pkg/homeassistant"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/states"
)

var get = sync.OnceValue(func() *sharedha.Controller {
	c := sharedha.New(sharedha.Settings{Home: config.HomeSection, API: config.APISection})
	c.FollowIntegration(states.Get())
	return c
})

func Get() *sharedha.Controller { return get() }

func init() { component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(5)) }
