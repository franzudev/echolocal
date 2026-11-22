package component

import (
	setting "github.com/franzudev/libcountertop/pkg/settings"

	"github.com/franzudev/libcountertop/pkg/hook"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
)

var Settings setting.Registry

var Reconnect hook.Hook[struct{}]
var Fire hook.Hook[sharedcomponent.Event]
var Subscribed hook.Hook[struct{}]
