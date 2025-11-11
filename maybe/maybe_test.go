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

func TestMaybeMap(t *testing.T) {
	x := m.Some(1)
	f := func(v int) string {
		if v == 1 {
			return "one"
		}
		return "not one"
	}

	z := m.Map(x, f)

	if z.Unwrap() != "one" {
		t.Fatalf("wanted Some(one), but got %v", z)
	}
}

func TestUnwrapOr(t *testing.T) {
	tests := []struct {
		name  string
		value m.Maybe[int]
		def   int
		want  int
	}{
		{name: "test some", value: m.Some(1), def: 2, want: 1},
		{name: "test none", value: m.None[int](), def: 4, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.UnwrapOr(tt.def)
			if got != tt.want {
				t.Fatalf("want %d, got %d", tt.want, got)
			}
		})
	}
}
