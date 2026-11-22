package all

import (
	"os"
	"os/exec"
	"slices"
	"strings"
	"testing"
)

var screenOnly = []string{
	"github.com/franzudev/libcountertop/pkg/display/panel",
	"github.com/franzudev/echolocal/internal/hardware/display",
	"github.com/franzudev/echolocal/internal/hardware/gpu",
	"github.com/franzudev/libcountertop/pkg/display/gpu",
	"github.com/franzudev/echolocal/internal/feature/visuals",
	"github.com/franzudev/echolocal/internal/feature/idle",
	"github.com/franzudev/echolocal/internal/feature/apps",
	"github.com/franzudev/libcountertop/pkg/display/apps",
	"github.com/franzudev/echolocal/internal/ui/visual",
	"github.com/franzudev/echolocal/internal/hardware/touch",
	"github.com/franzudev/echolocal/internal/ui/reveal",
	"github.com/franzudev/echolocal/internal/feature/brightness",
	"github.com/franzudev/echolocal/internal/hardware/alsps",
	"github.com/franzudev/echolocal/internal/feature/gui",
	"github.com/franzudev/echolocal/internal/feature/screen",
	"github.com/franzudev/echolocal/internal/feature/dashboard",
	"github.com/franzudev/echolocal/internal/feature/poster",
	"github.com/franzudev/echolocal/internal/feature/message",
	"github.com/franzudev/echolocal/internal/feature/assistant",
	"github.com/franzudev/echolocal/internal/feature/homecontrol",
	"github.com/franzudev/echolocal/internal/feature/a2dp",
	"github.com/franzudev/echolocal/internal/feature/drm",
	"github.com/franzudev/echolocal/internal/feature/weather",
	"github.com/franzudev/libcountertop/pkg/bluetooth",
	"github.com/franzudev/libcountertop/pkg/display/videopage",
	"github.com/franzudev/libcountertop/pkg/display/screens",
	"github.com/franzudev/libcountertop/pkg/display/gogui",
	"github.com/go-gui-org/go-gui/gui",
}

var castAudio = []string{
	"github.com/franzudev/echolocal/internal/feature/chromecast",
	"github.com/franzudev/echolocal/internal/feature/playout",
	"github.com/franzudev/libcountertop/pkg/media/cast",
	"github.com/franzudev/libcountertop/pkg/media/playout",
}

const dotYouTube = "github.com/franzudev/libcountertop/pkg/media/cast/protocols/youtube"

var dotCast = []string{
	dotYouTube,
	"github.com/franzudev/libcountertop/pkg/media/webm",
	"github.com/franzudev/libcountertop/pkg/media/hls",
}

var castOnly = []string{
	"github.com/franzudev/libcountertop/pkg/display/video",
	"github.com/franzudev/libcountertop/pkg/display/playoutscreen",
	"github.com/franzudev/libcountertop/pkg/media/cast/protocols/webreceiver",
	"github.com/franzudev/libcountertop/pkg/media/sealed",
	"github.com/franzudev/libcountertop/pkg/media/videostream",
}

var cameraOnly = []string{
	"github.com/franzudev/echolocal/internal/feature/livecam",
	"github.com/franzudev/echolocal/internal/feature/rtspd",
	"github.com/franzudev/echolocal/internal/feature/vision",
	"github.com/franzudev/echolocal/internal/hardware/mtkcamera",
	"github.com/franzudev/libcountertop/pkg/camera/rtspd",
	"github.com/franzudev/libcountertop/pkg/media/rtsp",
	"github.com/franzudev/libcountertop/pkg/camera/onvif",
}

func goList(t *testing.T, tags, pkg string, flags ...string) []string {
	t.Helper()
	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("no go toolchain to list the build with")
	}
	args := append([]string{"list", "-e"}, flags...)
	if tags != "" {
		args = append(args, "-tags", tags)
	}
	cmd := exec.Command(gobin, append(args, pkg)...)
	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=arm", "CGO_ENABLED=0")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("go list %s: %v", tags, err)
	}
	return strings.Fields(string(out))
}

func echodDeps(t *testing.T, tags string) []string {
	return goList(t, tags, "github.com/franzudev/echolocal/cmd/echod", "-deps")
}

func TestOnlyHelperBoardsEmbedTheHelper(t *testing.T) {
	for tags, want := range map[string]bool{
		"board_biscuit":          false,
		"board_checkers":         false,
		"board_checkers,payload": true,
		"board_cronos,payload":   true,
		"board_crown,payload":    true,
	} {
		files := goList(
			t,
			tags,
			"github.com/franzudev/echolocal/internal/parts",
			"-f",
			`{{join .EmbedPatterns " "}}`,
		)
		for _, part := range []string{"payload/echolocal-surface", "payload/echolocal-camera", "payload/libecholocal-camshim.so"} {
			if got := slices.Contains(files, part); got != want {
				t.Errorf("%q embeds %s %t, want %t (%v)", tags, part, got, want, files)
			}
		}
	}
}

func TestBiscuitLinksNoShowFeatures(t *testing.T) {
	for _, tag := range []string{"board_biscuit"} {
		deps := echodDeps(t, tag)
		for _, pkg := range slices.Concat(screenOnly, castOnly, cameraOnly) {
			if slices.Contains(deps, pkg) {
				t.Errorf("%q links Show feature %s", tag, pkg)
			}
		}
		for _, pkg := range deps {
			if showTree(pkg) {
				t.Errorf("%q links Show package %s", tag, pkg)
			}
		}
		for _, pkg := range slices.Concat(castAudio, dotCast) {
			if !slices.Contains(deps, pkg) {
				t.Errorf("%q does not link %s", tag, pkg)
			}
		}
	}
}

var (
	showTrees = []string{
		"github.com/franzudev/libcountertop/pkg/display/",
		"github.com/franzudev/libcountertop/pkg/camera/",
		"github.com/franzudev/libcountertop/pkg/input/touch",
	}
	biscuitDisplay = []string{
		"github.com/franzudev/libcountertop/pkg/display/theme",
	}
)

func showTree(pkg string) bool {
	if slices.Contains(biscuitDisplay, pkg) {
		return false
	}
	return slices.ContainsFunc(showTrees, func(tree string) bool {
		return strings.HasPrefix(pkg, tree)
	})
}

func TestShowsLinkCastAndCamera(t *testing.T) {
	for _, tag := range []string{"board_checkers", "board_cronos", "board_crown"} {
		deps := echodDeps(t, tag)
		for _, pkg := range slices.Concat(castAudio, castOnly, cameraOnly) {
			if !slices.Contains(deps, pkg) {
				t.Errorf("%s does not link %s", tag, pkg)
			}
		}
		if slices.Contains(deps, dotYouTube) {
			t.Errorf("%s links the Echo Dot's YouTube receiver", tag)
		}
	}
}

func TestScreenBoardsLinkTheScreen(t *testing.T) {
	for _, tag := range []string{
		"board_checkers",
		"board_cronos",
		"board_crown",
		"board_rook",
	} {
		deps := echodDeps(t, tag)
		for _, pkg := range screenOnly {
			if !slices.Contains(deps, pkg) {
				t.Errorf("%s does not link %s", tag, pkg)
			}
		}
	}
}
