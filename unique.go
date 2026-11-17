package seedfaker

import "fmt"

// defaultTries bounds the retry loops in UniqueN and Valid before they report that
// the value space is exhausted, instead of looping forever as fakerphp does.
const defaultTries = 1000

// UniqueN returns n distinct values from gen. It is deterministic: the stream is
// fixed, so the set of values is too. It returns the values gathered so far and an
// error once gen yields defaultTries duplicates in a row (the space is exhausted).
func UniqueN[T comparable](f *Faker, n int, gen func(*Faker) T) ([]T, error) {
	out := make([]T, 0, n)
	seen := make(map[T]struct{}, n)
	for len(out) < n {
		got := false
		for i := 0; i < defaultTries; i++ {
			v := gen(f)
			if _, dup := seen[v]; !dup {
				seen[v] = struct{}{}
				out = append(out, v)
				got = true
				break
			}
		}
		if !got {
			return out, fmt.Errorf("seedfaker: UniqueN exhausted at %d of %d", len(out), n)
		}
	}
	return out, nil
}

// Optional returns gen's value with probability p (clamped to [0,1]) and the zero
// value otherwise.
func Optional[T any](f *Faker, p float64, gen func(*Faker) T) T {
	var zero T
	if f.Float64() < p {
		return gen(f)
	}
	return zero
}

// Valid returns the first value from gen that ok accepts, or an error after
// defaultTries rejections.
func Valid[T any](f *Faker, gen func(*Faker) T, ok func(T) bool) (T, error) {
	var zero T
	for i := 0; i < defaultTries; i++ {
		if v := gen(f); ok(v) {
			return v, nil
		}
	}
	return zero, fmt.Errorf("seedfaker: Valid found no acceptable value in %d tries", defaultTries)
}
