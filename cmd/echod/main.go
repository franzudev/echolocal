package main

import (
	"github.com/franzudev/libcountertop/pkg/runtime/payload"

	"github.com/franzudev/echolocal/internal/cli/echod"
)

func main() {
	payload.Enter()
	echod.Execute()
}
