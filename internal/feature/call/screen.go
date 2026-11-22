//go:build !board_biscuit

package call

import (
	"github.com/franzudev/libcountertop/pkg/display/callvideo"
	"github.com/franzudev/libcountertop/pkg/display/callview"
	sharedcall "github.com/franzudev/libcountertop/pkg/media/call"

	"github.com/franzudev/echolocal/internal/feature/livecam"
	"github.com/franzudev/echolocal/internal/feature/message"
	"github.com/franzudev/echolocal/internal/feature/privacy"
	"github.com/franzudev/echolocal/internal/feature/shell"
	"github.com/franzudev/echolocal/internal/hardware/display"
)

func configureSettings(c *sharedcall.Calls) {}

func front(d *sharedcall.Dependencies) {
	d.Privacy = privacy.Get()
	d.Stage = callview.NewStage(shell.Get())
	d.Video = callvideo.New(callvideo.Dependencies{
		Display: display.Get(),
		Camera:  callvideo.NewCamera(livecam.Get().Sessions()),
	})
	d.Messages = message.Get()
}
