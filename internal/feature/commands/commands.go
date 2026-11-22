package commands

import (
	"sync"

	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/voice"

	"github.com/franzudev/echolocal/internal/component"
	"github.com/franzudev/echolocal/internal/config"
	"github.com/franzudev/echolocal/internal/feature/detect"
	"github.com/franzudev/echolocal/internal/feature/discovery"
	"github.com/franzudev/echolocal/internal/feature/homeassistant"
	"github.com/franzudev/echolocal/internal/hardware/mic"
	"github.com/franzudev/echolocal/internal/layout"
	"github.com/franzudev/echolocal/internal/parts"
)

func init() {
	component.Register(sharedcomponent.Device, Get, sharedcomponent.Order(41))
}

type Commands struct{ *voice.Listener }

func (c *Commands) Restore(config.Config) { c.Apply() }

func (c *Commands) Close() error {
	c.Stop()
	return nil
}

var get = sync.OnceValue(func() *Commands {
	c := &Commands{voice.NewListener(voice.Options{
		Helper:   layout.SpeechHelper,
		Binary:   parts.Speech(),
		ModelDir: layout.ModelDir,
		Source:   mic.Get(),
		Paused:   detect.Get().Paused,
		Settings: config.VoiceSection,
		Home:     homeassistant.Get(),
		Peers:    discovery.Get().PeerNames,
		Retired: []string{
			layout.StateDir + "/bin/echolocal-voice",
			layout.ModelDir + "/whistle-*.cact",
		},
	})}
	return c
})

func Get() *Commands { return get() }
