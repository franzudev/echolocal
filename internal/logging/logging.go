package logging

import (
	"github.com/franzudev/libcountertop/pkg/android/logd"

	"github.com/franzudev/echolocal/internal/layout"
)

var Log = logd.NewLog(layout.LogTag)
