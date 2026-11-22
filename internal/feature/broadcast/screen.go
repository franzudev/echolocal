//go:build !board_biscuit

package broadcast

import (
	shared "github.com/franzudev/libcountertop/pkg/media/broadcast"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/message"
)

func messages() shared.Messages {
	if !component.Board().Has(board.Panel) {
		return nil
	}
	return message.Get()
}
