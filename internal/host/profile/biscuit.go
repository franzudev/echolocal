package profile

import (
	"github.com/franzudev/echolocal/internal/android/services"
	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/host/bootimg"
	"github.com/franzudev/echolocal/internal/host/sysimg"
)

// biscuit is the 2nd-generation Echo Dot.
var biscuit = Profile{
	Board:  board.Biscuit,
	Boot:   bootimg.Biscuit,
	System: sysimg.Biscuit,

	Disable: services.Disabled,
	Enable:  services.Enabled,
}
