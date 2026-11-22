package display

import (
	"testing"

	"github.com/franzudev/libcountertop/pkg/display/geometry"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
)

func TestTheShowsAreMountedLandscapeInverted(t *testing.T) {
	was := component.Board()
	defer component.Default().Use(was)

	for _, b := range []board.Board{board.Checkers, board.Cronos} {
		component.Default().Use(b)
		if Mounted() != geometry.Rotate270 {
			t.Errorf("%s is mounted %v, want %v", b.Codename, Mounted(), geometry.Rotate270)
		}
	}
}
