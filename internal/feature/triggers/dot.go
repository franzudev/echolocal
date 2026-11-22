//go:build board_biscuit

package triggers

import setting "github.com/franzudev/libcountertop/pkg/settings"

func commandNames() []string { return nil }

func applyCommands() {}

func labelChoices() func() []setting.Option { return nil }
