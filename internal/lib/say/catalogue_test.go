package say_test

import (
	"testing"

	"github.com/franzudev/libcountertop/pkg/say/saytest"
)

func TestEveryIdentifierTheCodeAsksForExists(t *testing.T) { saytest.Catalogue(t, "../../..") }
