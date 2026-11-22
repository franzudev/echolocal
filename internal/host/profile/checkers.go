package profile

import (
	"github.com/franzudev/echolocal/internal/android/services"
	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/host/bootimg"
	"github.com/franzudev/echolocal/internal/host/sysimg"
)

// checkers is the 1st-generation Echo Show 5.
var checkers = Profile{
	Board:  board.Checkers,
	Boot:   bootimg.Checkers,
	System: sysimg.Checkers,

	Disable: append([]string{
		"dha_service",
		"integrity_logger",
		"amazonfiled",
		"meshmgrservice",
		"inlocservice",
		"shblemeshd",
		"shchipd",
		"shlocalskillsd",
		"smarthomewifid",
		"securetime",
		"halo-hub-broker",
		"wifi_log_levels",
		"aal",
		"aal-metric",
		"touchlogger",
		"installd",
		"gatekeeperd",
		"keystore",
		"audioserver",
	}, services.Disabled...),

	Enable: append([]string{
		"media",
		"mediacodec",
		"mediadrm",
		"mediaextractor",
		"cameraserver",
	}, services.Enabled...),
}
