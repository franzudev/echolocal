//go:build !board_biscuit

package videoplayer

import (
	sharedvolume "github.com/franzudev/libcountertop/pkg/audio/volume"
	sharedpage "github.com/franzudev/libcountertop/pkg/display/videopage"

	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/feature/volume"
)

func NewPage(c sharedpage.Controls) *sharedpage.Page {
	return sharedpage.New(
		c,
		sharedpage.Dependencies{
			Shell:  shell.Get(),
			Volume: sharedvolume.For(volume.Get(), config.StreamMedia),
			Left:   func() { media.Card().Left() },
		},
	)
}
