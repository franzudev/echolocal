package echod

import (
	"github.com/spf13/cobra"
	harness "github.com/franzudev/libcountertop/pkg/runtime/control"

	"github.com/franzudev/echolocal/internal/feature/control"
)

func newCtlCmd() *cobra.Command {
	return harness.ClientCommand(harness.ClientOptions{Name: "echod", Socket: control.Socket})
}
