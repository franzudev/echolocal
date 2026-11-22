//go:build !board_biscuit

package actions

import "github.com/franzudev/echolocal/internal/feature/tv"

func (a *Actions) addTV() { a.reg.Add(tv.Get().VoiceAction()) }
