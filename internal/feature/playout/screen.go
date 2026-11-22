//go:build !board_biscuit

package playout

import (
	"context"

	"github.com/franzudev/libcountertop/pkg/display/geometry"
	"github.com/franzudev/libcountertop/pkg/display/playoutscreen"
	"github.com/franzudev/libcountertop/pkg/display/video"
	"github.com/franzudev/libcountertop/pkg/media/playout"
	"github.com/franzudev/libcountertop/pkg/media/videostream"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/feature/videoplayer"
	"github.com/franzudev/echolocal/internal/hardware/display"
)

func screen() playout.Screen {
	return playoutscreen.New(playoutscreen.Options{
		NewPage: videoplayer.NewPage,
		NewUnder: func() *videostream.Beneath {
			return &videostream.Beneath{
				Rotation: func() geometry.Orientation { return display.Get().Orientation() },
				Size:     func() (int, int) { return display.Get().Native() },
			}
		},
		Video: func(ctx context.Context, b *videostream.Beneath, s videostream.Stream) (videostream.Report, error) {
			player := video.Player{
				Helper:         display.Get().Helper(),
				SecureDecoders: component.Board().SecureDecoders,
			}
			return player.On(ctx, b, s)
		},
	})
}
