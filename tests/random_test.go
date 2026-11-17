package tests

import (
	"strings"
	"testing"

	"github.com/saphemmy/seedfaker"
)

func assertPanics(t *testing.T, name string, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("%s: expected panic, got none", name)
		}
	}()
	fn()
}

func TestIntNBounds(t *testing.T) {
	f := seedfaker.New(1)
	for range 10000 {
		if v := f.IntN(10); v < 0 || v >= 10 {
			t.Fatalf("IntN(10) = %d, out of [0,10)", v)
		}
	}
}

func TestIntNPanicsOnNonPositive(t *testing.T) {
	assertPanics(t, "IntN(0)", func() { seedfaker.New(1).IntN(0) })
	assertPanics(t, "IntN(-5)", func() { seedfaker.New(1).IntN(-5) })
}

func TestIntBetween(t *testing.T) {
	f := seedfaker.New(2)
	for range 10000 {
		if v := f.IntBetween(5, 7); v < 5 || v > 7 {
			t.Fatalf("IntBetween(5,7) = %d", v)
		}
	}
	assertPanics(t, "IntBetween(7,5)", func() { f.IntBetween(7, 5) })
	if v := f.IntBetween(4, 4); v != 4 {
		t.Fatalf("IntBetween(4,4) = %d", v)
	}
}

func TestFloat64Range(t *testing.T) {
	f := seedfaker.New(3)
	for range 10000 {
		if v := f.Float64(); v < 0 || v >= 1 {
			t.Fatalf("Float64 = %v, out of [0,1)", v)
		}
	}
}

func TestFloatBetween(t *testing.T) {
	f := seedfaker.New(3)
	for range 10000 {
		if v := f.FloatBetween(-2, 2); v < -2 || v >= 2 {
			t.Fatalf("FloatBetween(-2,2) = %v", v)
		}
	}
	assertPanics(t, "FloatBetween(2,-2)", func() { f.FloatBetween(2, -2) })
}

func TestNumber(t *testing.T) {
	f := seedfaker.New(4)
	for range 2000 {
		if v := f.Number(3); v < 0 || v > 999 {
			t.Fatalf("Number(3) = %d", v)
		}
	}
	assertPanics(t, "Number(0)", func() { f.Number(0) })
	assertPanics(t, "Number(19)", func() { f.Number(19) })
}

func TestDigitHelpers(t *testing.T) {
	f := seedfaker.New(5)
	for range 2000 {
		if d := f.DigitNot(3); d == 3 || d < 0 || d > 9 {
			t.Fatalf("DigitNot(3) = %d", d)
		}
		if d := f.DigitNotNull(); d < 1 || d > 9 {
			t.Fatalf("DigitNotNull = %d", d)
		}
	}
	assertPanics(t, "DigitNot(10)", func() { f.DigitNot(10) })
}

func TestFormatFillers(t *testing.T) {
	f := seedfaker.New(6)
	if n := f.Numerify("###"); len(n) != 3 || strings.IndexFunc(n, func(r rune) bool { return r < '0' || r > '9' }) != -1 {
		t.Fatalf("Numerify(###) = %q", n)
	}
	if l := f.Lexify("????"); len(l) != 4 || strings.IndexFunc(l, func(r rune) bool { return r < 'A' || r > 'Z' }) != -1 {
		t.Fatalf("Lexify(????) = %q", l)
	}
	if b := f.Bothify("prefix-#?"); !strings.HasPrefix(b, "prefix-") || len(b) != 9 {
		t.Fatalf("Bothify = %q", b)
	}
}

func TestBool(t *testing.T) {
	f := seedfaker.New(7)
	trues := 0
	for range 10000 {
		if f.Bool() {
			trues++
		}
	}
	if trues < 4000 || trues > 6000 {
		t.Fatalf("Bool is badly skewed: %d/10000 true", trues)
	}
}

func TestPickAndShuffle(t *testing.T) {
	f := seedfaker.New(8)
	s := []int{1, 2, 3, 4, 5}
	for range 1000 {
		if v := seedfaker.Pick(f, s); v < 1 || v > 5 {
			t.Fatalf("Pick = %d", v)
		}
	}
	assertPanics(t, "Pick(empty)", func() { seedfaker.Pick(f, []int{}) })

	cp := append([]int(nil), s...)
	seedfaker.Shuffle(f, cp)
	sum := 0
	for _, v := range cp {
		sum += v
	}
	if sum != 15 {
		t.Fatalf("Shuffle lost elements: %v", cp)
	}
}

func TestWeighted(t *testing.T) {
	f := seedfaker.New(9)
	s := []string{"a", "b"}
	counts := map[string]int{}
	for range 10000 {
		counts[seedfaker.Weighted(f, s, []int{9, 1})]++
	}
	if counts["a"] <= counts["b"] {
		t.Fatalf("Weighted ignored weights: %v", counts)
	}
	assertPanics(t, "Weighted mismatch", func() { seedfaker.Weighted(f, s, []int{1}) })
	assertPanics(t, "Weighted zero", func() { seedfaker.Weighted(f, s, []int{0, 0}) })
}
