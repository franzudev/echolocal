package dashboard

import (
	sharedlib "github.com/franzudev/libcountertop/pkg/display/dashboard"

	"github.com/franzudev/echolocal/internal/feature/shell"
)

var strip = sharedlib.NewTabs(func() { shell.Get().Redraw() })

func Tabs() *sharedlib.Tabs { return strip }
