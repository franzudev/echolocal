package splash

import (
	bootview "github.com/franzudev/libcountertop/pkg/display/boot"
	"github.com/franzudev/libcountertop/pkg/display/panel"
	sharedui "github.com/franzudev/libcountertop/pkg/display/ui"

	"github.com/franzudev/echolocal/internal/hardware/display"
	"github.com/franzudev/echolocal/internal/layout"
	"github.com/franzudev/echolocal/internal/ui"
)

func init() { display.Get().Stranded(drawStranded) }

func drawStranded(p *panel.Panel) error {
	bootview.Stranded(sharedui.Of(p), chosen(), ui.Logo().Draw, layout.Version)
	return nil
}
