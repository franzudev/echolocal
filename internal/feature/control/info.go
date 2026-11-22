package control

import (
	"time"

	harness "github.com/franzudev/libcountertop/pkg/runtime/control"

	"github.com/franzudev/libcountertop/pkg/system/metrics"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/layout"
)

func deviceInfo() harness.DeviceInfo {
	d := config.Get().Device
	return harness.DeviceInfo{
		Version: layout.Version,
		Commit:  layout.GitCommit,
		Built:   layout.BuildDate,
		Board:   component.Board().Codename,
		Name:    d.Name,
		Address: d.Addr,
		Model:   d.Board.Model,
		Uptime:  time.Duration(metrics.Reader{}.Uptime().Value * float64(time.Second)),
	}
}
