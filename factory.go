package seedfaker

// Factory builds deterministic instances of T from a Faker. It is the ORM-agnostic
// way to seed: Build produces one value and Create persists each through a callback
// you supply, so the core depends on no database or ORM.
type Factory[T any] struct {
	// Build produces one instance. It should draw every field from f so the result
	// is deterministic.
	Build func(f *Faker) T
}

// MakeOne builds a single instance.
func (fac Factory[T]) MakeOne(f *Faker) T { return fac.Build(f) }

// Make builds n instances; f advances deterministically between them.
func (fac Factory[T]) Make(f *Faker, n int) []T {
	out := make([]T, n)
	for i := range out {
		out[i] = fac.Build(f)
	}
	return out
}

// Create builds n instances and passes each to persist (a GORM, ent, bun, sqlx or
// raw-SQL insert). It stops at the first persist error, returning the instances
// stored so far and that error.
func (fac Factory[T]) Create(f *Faker, n int, persist func(T) error) ([]T, error) {
	out := make([]T, 0, n)
	for i := 0; i < n; i++ {
		v := fac.Build(f)
		if err := persist(v); err != nil {
			return out, err
		}
		out = append(out, v)
	}
	return out, nil
}
