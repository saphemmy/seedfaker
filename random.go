package seedfaker

import "math/bits"

// IntN returns a uniform int in [0, n); it panics if n <= 0. Every generator is
// built on IntN, whose Lemire reduction over raw PCG output keeps results stable
// across Go versions.
func (f *Faker) IntN(n int) int {
	if n <= 0 {
		panic("seedfaker: IntN requires n > 0")
	}
	return int(f.u64n(uint64(n)))
}

// u64n returns a uniform uint64 in [0, n) (Lemire with rejection); n must be > 0.
func (f *Faker) u64n(n uint64) uint64 {
	x := f.src.Uint64()
	hi, lo := bits.Mul64(x, n)
	if lo < n {
		thresh := -n % n
		for lo < thresh {
			x = f.src.Uint64()
			hi, lo = bits.Mul64(x, n)
		}
	}
	return hi
}

// Uint64 returns the next raw value; use it to build custom generators.
func (f *Faker) Uint64() uint64 { return f.src.Uint64() }

// Bool reports a random boolean.
func (f *Faker) Bool() bool { return f.src.Uint64()&1 == 0 }

// IntBetween returns a uniform int in [lo, hi]; it panics if lo > hi.
func (f *Faker) IntBetween(lo, hi int) int {
	if lo > hi {
		panic("seedfaker: IntBetween requires lo <= hi")
	}
	return lo + f.IntN(hi-lo+1)
}

// Digit returns a decimal digit 0-9.
func (f *Faker) Digit() int { return f.IntN(10) }
