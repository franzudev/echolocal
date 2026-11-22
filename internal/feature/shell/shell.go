package shell

import (
	sharedshell "github.com/franzudev/libcountertop/pkg/display/shell"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"

	"github.com/franzudev/echolocal/internal/component"
)

var shell = sharedshell.New()

func Get() *sharedshell.Shell { return shell }

func init() { component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(34)) }
