package states

import (
	sharedstates "github.com/franzudev/libcountertop/pkg/homeassistant/states"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
)

type States = sharedstates.States
type Value = sharedstates.Value

var shared = sharedstates.New()

func Get() *States { return shared }
func init()        { component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(5)) }
