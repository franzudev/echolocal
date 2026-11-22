//go:build !board_biscuit

package all

import (
	_ "github.com/franzudev/echolocal/internal/feature/livecam"
	_ "github.com/franzudev/echolocal/internal/feature/rtspd"
	_ "github.com/franzudev/echolocal/internal/feature/vision"
)
