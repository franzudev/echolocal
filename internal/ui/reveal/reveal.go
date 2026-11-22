package reveal

import (
	_ "embed"

	bootview "github.com/franzudev/libcountertop/pkg/display/boot"
	"github.com/franzudev/libcountertop/pkg/display/theme"

	"github.com/franzudev/echolocal/internal/ui"
)

//go:embed reveal.glsl
var revealShader string

func New(palette theme.Theme) *bootview.Reveal {
	return bootview.NewReveal(palette, revealShader, ui.Logo().On(palette.Background), 768, 512)
}
