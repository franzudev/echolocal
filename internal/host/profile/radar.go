package profile

import (
	"github.com/franzudev/echolocal/internal/android/services"
	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/host/bootimg"
	"github.com/franzudev/echolocal/internal/host/sysimg"
)

var radar = Profile{
	Board:  board.Radar,
	Boot:   bootimg.Radar,
	System: sysimg.Radar,

	Disable: services.Disabled,
	Enable:  services.Enabled,
}
