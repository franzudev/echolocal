//go:build !board_biscuit

package alerts

import (
	shared "github.com/franzudev/libcountertop/pkg/audio/alerts"
)

func configure(a *shared.Alerts) {}
