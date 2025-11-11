package maybe_test

import (
	"testing"

	m "github.com/elmq0022/maybe/maybe"
)

func TestMaybeSome(t *testing.T) {
	x := m.Some(1)

	if !x.IsSome() {
		t.Fatalf("wanted x.isSome() == True but got %t", x.IsSome())
	}
}

func TestMaybeNone(t *testing.T) {
	y := m.None[int]()
	if !y.IsNone() {
		t.Fatalf("wanted y.IsNone() == True, but got %t", y.IsNone())
	}
}
