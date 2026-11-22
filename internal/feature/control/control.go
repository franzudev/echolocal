package control

import (
	"sync"
	"time"

	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	harness "github.com/franzudev/libcountertop/pkg/runtime/control"
	"github.com/franzudev/libcountertop/pkg/runtime/service"

	"github.com/franzudev/echolocal/internal/component"
)

const Socket = "@echod"

type Control struct {
	*harness.Server

	inputOnce sync.Once
	input     *harness.Input
}

func init() {
	component.Register(
		sharedcomponent.Device,
		Get,
		sharedcomponent.Order(90),
		sharedcomponent.Supervise(service.Restart(time.Second, time.Minute)),
	)
}

var Get = sync.OnceValue(func() *Control { return build(Socket) })

func build(addr string) *Control {
	c := &Control{}
	c.Server = harness.NewServer(addr, harness.Executes(c.tree))
	return c
}

func (c *Control) Name() string { return "control" }
