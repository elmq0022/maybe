package maybe_test

import (
	"testing"

	m "github.com/elmq0022/maybe/maybe"
)

func TestMaybeInt(t *testing.T) {
	x := m.Some[int](1)

	if !x.IsSome() {
		t.Fatalf("wanted x.isSome() to be True but got %t", x.IsSome())
	}
}
