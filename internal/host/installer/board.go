package installer

import (
	"github.com/franzudev/libcountertop/pkg/host/adb"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/host/sysimg"
)

func BoardOf(d *adb.Device) board.Board {
	device, err := d.Getprop("ro.product.device")
	if err != nil {
		return board.Biscuit
	}
	if b, ok := board.For(device); ok {
		return b
	}
	return board.Biscuit
}

func (r *run) board() board.Board {
	if r.on.Device == "" {
		if r.on = r.cfg.Profile.Board; r.on.Device == "" {
			r.on = BoardOf(r.d)
		}
	}
	return r.on
}

func (r *run) system() sysimg.Layout { return r.cfg.Profile.System }
