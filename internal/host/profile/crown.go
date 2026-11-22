package profile

import (
	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/host/bootimg"
	"github.com/franzudev/echolocal/internal/host/sysimg"
)

var crown = Profile{
	Board:  board.Crown,
	Boot:   bootimg.Crown,
	System: sysimg.Crown,

	Disable: checkers.Disable,
	Enable:  checkers.Enable,
}
