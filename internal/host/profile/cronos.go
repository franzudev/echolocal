package profile

import (
	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/host/bootimg"
	"github.com/franzudev/echolocal/internal/host/sysimg"
)

var cronos = Profile{
	Board:  board.Cronos,
	Boot:   bootimg.Cronos,
	System: sysimg.Cronos,

	Disable: checkers.Disable,
	Enable:  checkers.Enable,
}
