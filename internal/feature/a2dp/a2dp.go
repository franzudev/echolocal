//go:build !board_biscuit

package a2dp

import (
	"sync"
	"time"

	sharedvolume "github.com/franzudev/libcountertop/pkg/audio/volume"
	sharedspeaker "github.com/franzudev/libcountertop/pkg/bluetooth/a2dp"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/service"

	"github.com/franzudev/echolocal/internal/board"
	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/media"
	"github.com/franzudev/echolocal/internal/feature/volume"
	"github.com/franzudev/echolocal/internal/hardware/ble"
	"github.com/franzudev/echolocal/internal/hardware/speaker"
	"github.com/franzudev/echolocal/internal/layout"
)

var once sync.Once
var shared *sharedspeaker.Sink

func build() *sharedspeaker.Sink {
	return sharedspeaker.New(sharedspeaker.Dependencies{
		Queue:            speaker.Get().Lane(config.StreamMedia),
		Radio:            ble.Get,
		Bonds:            layout.StateDir + "/bonds",
		CaptureDirectory: "/data/local/tmp",
		Name:             func() string { return config.Get().Device.Name },
		Settings:         config.BluetoothSection,
		Volume:           sharedvolume.For(volume.Get(), config.StreamMedia),
		Media:            media.Get(),
		Arbitration:      speaker.Sound().Backgrounds(),
		Ducking:          config.MediaSection,
	})
}
func Get() *sharedspeaker.Sink {
	once.Do(func() {
		shared = build()
		component.Settings.Add(shared.Controls())
	})
	return shared
}
func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(70),
		component.Needs(board.Panel),
		sharedcomponent.Supervise(service.Restart(5*time.Second, time.Minute)),
	)
}
