package seedfaker

// Float64 returns a float uniformly in [0.0, 1.0), with 53 bits of precision.
func (f *Faker) Float64() float64 { return float64(f.src.Uint64()>>11) / (1 << 53) }

// FloatBetween returns a float uniformly in [min, max). It panics if min > max.
func (f *Faker) FloatBetween(min, max float64) float64 {
	if min > max {
		panic("seedfaker: FloatBetween requires min <= max")
	}
	return min + f.Float64()*(max-min)
}

// Number returns a non-negative integer with up to digits digits. It panics unless
// digits is in 1..18 (the range that always fits an int64).
func (f *Faker) Number(digits int) int {
	if digits < 1 || digits > 18 {
		panic("seedfaker: Number requires digits in 1..18")
	}
	max := 1
	for i := 0; i < digits; i++ {
		max *= 10
	}
	return f.IntN(max)
}

// DigitNot returns a decimal digit other than d; it panics unless d is in 0..9.
func (f *Faker) DigitNot(d int) int {
	if d < 0 || d > 9 {
		panic("seedfaker: DigitNot requires d in 0..9")
	}
	n := f.IntN(9)
	if n >= d {
		n++
	}
	return n
}

// DigitNotNull returns a decimal digit 1-9.
func (f *Faker) DigitNotNull() int { return f.IntN(9) + 1 }

// Letter returns an uppercase letter A-Z.
func (f *Faker) Letter() byte { return byte('A' + f.IntN(26)) }
