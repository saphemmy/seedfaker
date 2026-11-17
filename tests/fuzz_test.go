package tests

import (
	"strings"
	"testing"

	"github.com/saphemmy/seedfaker"
)

// FuzzIntN: for any seed and any positive n, IntN stays in [0, n) and never panics.
func FuzzIntN(f *testing.F) {
	f.Add(uint64(1), 10)
	f.Add(uint64(99), 1)
	f.Fuzz(func(t *testing.T, seed uint64, n int) {
		if n <= 0 {
			t.Skip()
		}
		if v := seedfaker.New(seed).IntN(n); v < 0 || v >= n {
			t.Fatalf("IntN(%d) = %d", n, v)
		}
	})
}

// FuzzBothify: any pattern expands without panic and leaves no unfilled tokens.
func FuzzBothify(f *testing.F) {
	f.Add("ID-##-??")
	f.Add("")
	f.Fuzz(func(t *testing.T, pattern string) {
		out := seedfaker.New(1).Bothify(pattern)
		if strings.ContainsAny(out, "#?*") {
			t.Fatalf("Bothify(%q) left tokens: %q", pattern, out)
		}
	})
}
