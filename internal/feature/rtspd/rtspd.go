package rtspd

import (
	"sync"

	esphome "github.com/ygelfand/go-esphome-device"

	"github.com/franzudev/libcountertop/pkg/camera/live"
	"github.com/franzudev/libcountertop/pkg/camera/onvif"
	camerartspd "github.com/franzudev/libcountertop/pkg/camera/rtspd"
	"github.com/franzudev/libcountertop/pkg/media/rtsp"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	setting "github.com/franzudev/libcountertop/pkg/settings"
	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/android/prop"
	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/livecam"
	"github.com/franzudev/echolocal/internal/feature/privacy"
	"github.com/franzudev/echolocal/internal/feature/vision"
	"github.com/franzudev/echolocal/internal/hardware/mic"
	"github.com/franzudev/echolocal/internal/layout"
)

const (
	Port      = 8554
	ONVIFPort = 8000
)

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(70),
		component.Needs(board.Camera),
	)
}

type Server struct {
	knobs *setting.Controls[schema.RTSP]
	srv   *camerartspd.Server
}

var (
	once   sync.Once
	shared *Server
)

func Get() *Server {
	once.Do(func() {
		shared = &Server{srv: camerartspd.New(camerartspd.Options{
			Port:     Port,
			Identity: rtsp.Identity{Name: "EchoLocal", CNAMEPrefix: "echolocal-"},
			Source:   livecam.Get().Sessions(),
			Helper:   "echolocal-camera",
			Served:   livecam.Get().Served,
			FPS:      live.FPS,
			Bitrate:  livecam.Get().Bitrate,
			ONVIF:    describe,
			Auth:     camerartspd.Credentials(func() schema.RTSP { return config.Get().RTSP }),
			Mic:      microphone{},
			MicMuted: privacy.Get().MicMuted,
			Audio:    func() bool { return config.Get().RTSP.Audio },
		})}
		shared.build()
		livecam.Get().StreamsChanged.Listen(func([]int) { shared.reload() })
	})
	return shared
}

type microphone struct{}

func (microphone) Frames() (<-chan []int16, func()) { return mic.Get().Listen("rtsp") }

func describe() onvif.Config {
	dev := config.Get().Device
	serial, _ := prop.Get("ro.serialno")
	if serial == "" {
		serial = dev.Name
	}
	return onvif.Config{
		Port: ONVIFPort,
		Device: onvif.Device{
			Manufacturer: layout.Manufacturer,
			Model:        dev.Board.Model,
			Firmware:     layout.Version,
			Serial:       serial,
			Hardware:     dev.Board.Model,
			Name:         dev.Name,
		},
		Hardware: "EchoLocal",
		Snapshot: vision.Get().Still,
	}
}

func (s *Server) Name() string { return "rtsp" }

func (s *Server) Entities() []esphome.Entity { return s.knobs.Entities() }

func (s *Server) Restore(cfg config.Config) {
	s.knobs.PublishFrom(cfg.RTSP)
	if cfg.RTSP.Enabled {
		s.srv.Start()
	}
}

func (s *Server) build() {
	s.knobs = camerartspd.Controls(
		component.DeviceCamera,
		func() schema.RTSP { return config.Get().RTSP },
		s.apply,
	)
	component.Settings.Add(s.knobs)
}

func (s *Server) Settings() *setting.Controls[schema.RTSP] { return s.knobs }

func (s *Server) apply(at schema.RTSP) error {
	was := config.Get().RTSP
	err := config.Set().RTSP().All(at)
	switch {
	case !at.Enabled:
		s.srv.Stop()
	case was.Audio != at.Audio:
		s.srv.Stop()
		s.srv.Start()
	default:
		s.srv.Start()
	}
	return err
}

func (s *Server) reload() {
	if !config.Get().RTSP.Enabled {
		return
	}
	s.srv.Stop()
	s.srv.Start()
}
