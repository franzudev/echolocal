package ui

import (
	sharedui "github.com/franzudev/libcountertop/pkg/display/ui"

	"github.com/franzudev/echolocal/internal/brand"
)

var logo = sharedui.NewLogo(brand.LogoLight, brand.LogoDark)

func Logo() *sharedui.Logo { return logo }
