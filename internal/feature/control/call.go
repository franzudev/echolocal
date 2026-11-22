//go:build !board_biscuit

package control

import (
	"github.com/spf13/cobra"
	harness "github.com/franzudev/libcountertop/pkg/runtime/control"
	sharedcmd "github.com/franzudev/libcountertop/pkg/runtime/control/callcmd"

	"github.com/franzudev/echolocal/internal/feature/call"
	"github.com/franzudev/echolocal/internal/feature/discovery"
)

func callCommand() *cobra.Command {
	c := call.Get()
	return harness.CallCommand(
		harness.CallActions{
			Dial:   dial,
			State:  callState,
			Answer: c.Answer,
			Hangup: c.Hangup,
			Mute:   c.Mute,
			Camera: c.Camera,
		},
	)
}
func dial(a []string) (string, error) {
	return sharedcmd.Dial(call.Get(), discovery.Get().Peers(), a)
}
func callState(a []string) (string, error) { return sharedcmd.State(call.Get(), a) }
