package seedfaker

// Pick returns a random element of s; it panics if s is empty.
func Pick[T any](f *Faker, s []T) T {
	if len(s) == 0 {
		panic("seedfaker: Pick from empty slice")
	}
	return s[f.IntN(len(s))]
}

// Shuffle permutes s in place (Fisher-Yates).
func Shuffle[T any](f *Faker, s []T) {
	for i := len(s) - 1; i > 0; i-- {
		j := f.IntN(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

// Weighted returns an element of s chosen in proportion to w, where w[i] is the
// relative weight of s[i]. It panics if the slices differ in length, are empty,
// hold a negative weight, or sum to zero.
func Weighted[T any](f *Faker, s []T, w []int) T {
	if len(s) == 0 || len(s) != len(w) {
		panic("seedfaker: Weighted needs equal, non-empty slices")
	}
	total := 0
	for _, x := range w {
		if x < 0 {
			panic("seedfaker: Weighted needs non-negative weights")
		}
		total += x
	}
	if total == 0 {
		panic("seedfaker: Weighted needs a positive total weight")
	}
	r := f.IntN(total)
	for i, x := range w {
		if r < x {
			return s[i]
		}
		r -= x
	}
	return s[len(s)-1]
}

func (f *Faker) pick(s []string) string { return s[f.IntN(len(s))] }
