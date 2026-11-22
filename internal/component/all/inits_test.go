package all

import (
	"testing"

	"github.com/franzudev/libcountertop/pkg/build/inits"
)

func TestNoScatteredInits(t *testing.T) {
	found, err := inits.Scattered("../../..")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range found {
		t.Error(f)
	}
}
