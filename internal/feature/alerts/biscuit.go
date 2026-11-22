//go:build board_biscuit

package alerts

import (
	shared "github.com/franzudev/libcountertop/pkg/audio/alerts"

	"github.com/franzudev/echolocal/internal/feature/tones"
)

func configure(a *shared.Alerts) { a.Table = shared.TimerTable(tones.Get().Choices) }
