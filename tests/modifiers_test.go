package tests

import (
	"testing"

	"github.com/saphemmy/seedfaker"
)

func TestUniqueNReturnsDistinct(t *testing.T) {
	f := seedfaker.New(1)
	got, err := seedfaker.UniqueN(f, 10, func(f *seedfaker.Faker) int { return f.IntN(10) })
	if err != nil {
		t.Fatalf("UniqueN(10 of 10): %v", err)
	}
	seen := map[int]bool{}
	for _, v := range got {
		if seen[v] {
			t.Fatalf("duplicate %d in unique set %v", v, got)
		}
		seen[v] = true
	}
	if len(got) != 10 {
		t.Fatalf("got %d values, want 10", len(got))
	}
}

func TestUniqueNExhaustionReturnsPartialAndError(t *testing.T) {
	f := seedfaker.New(2)
	// Only 10 distinct digits exist, so asking for 11 must fail after yielding 10.
	got, err := seedfaker.UniqueN(f, 11, func(f *seedfaker.Faker) int { return f.IntN(10) })
	if err == nil {
		t.Fatal("expected exhaustion error asking for 11 of 10")
	}
	if len(got) != 10 {
		t.Fatalf("partial result = %d values, want 10", len(got))
	}
}

func TestUniqueNZeroIsEmpty(t *testing.T) {
	f := seedfaker.New(3)
	got, err := seedfaker.UniqueN(f, 0, func(f *seedfaker.Faker) string { return f.Name() })
	if err != nil || len(got) != 0 {
		t.Fatalf("UniqueN(0) = %v, %v", got, err)
	}
}

func TestUniqueNIsDeterministic(t *testing.T) {
	gen := func(f *seedfaker.Faker) string { return f.Email() }
	a, _ := seedfaker.UniqueN(seedfaker.New(7), 100, gen)
	b, _ := seedfaker.UniqueN(seedfaker.New(7), 100, gen)
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("unique set diverged at %d: %q != %q", i, a[i], b[i])
		}
	}
}

func TestOptionalBounds(t *testing.T) {
	gen := func(f *seedfaker.Faker) int { return f.IntN(100) + 1 } // never zero
	f := seedfaker.New(4)
	for range 1000 {
		if v := seedfaker.Optional(f, 1.0, gen); v == 0 {
			t.Fatal("Optional(p=1) returned the zero value")
		}
	}
	for range 1000 {
		if v := seedfaker.Optional(f, 0.0, gen); v != 0 {
			t.Fatalf("Optional(p=0) returned %d", v)
		}
	}
	// Out-of-range weights clamp to the nearest edge behaviour.
	for range 1000 {
		if v := seedfaker.Optional(f, -1, gen); v != 0 {
			t.Fatalf("Optional(p<0) returned %d", v)
		}
		if v := seedfaker.Optional(f, 2, gen); v == 0 {
			t.Fatal("Optional(p>1) returned the zero value")
		}
	}
}

func TestValidFiltersAndExhausts(t *testing.T) {
	even := func(v int) bool { return v%2 == 0 }
	f := seedfaker.New(5)
	for range 1000 {
		v, err := seedfaker.Valid(f, func(f *seedfaker.Faker) int { return f.IntN(100) }, even)
		if err != nil || !even(v) {
			t.Fatalf("Valid even: %d, %v", v, err)
		}
	}
	// A predicate nothing satisfies must error, not loop forever.
	_, err := seedfaker.Valid(f, func(f *seedfaker.Faker) int { return f.IntN(100) }, func(int) bool { return false })
	if err == nil {
		t.Fatal("expected exhaustion error from an impossible predicate")
	}
}
