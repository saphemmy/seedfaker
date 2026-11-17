package seedfaker

import (
	"math/rand/v2"
	"time"
)

// reference is the fixed "now" the date generators measure from, so Date, Past and
// Future are deterministic and never read the wall clock. Override with
// WithReferenceTime.
var reference = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

// Faker generates deterministic fake data. Build one with New; the zero value is
// unusable. A Faker is not safe for concurrent use — use Clone or Fork for parallel
// work.
type Faker struct {
	src *rand.PCG
	loc Locale
	ref time.Time
}

// New returns a Faker seeded for reproducible output. The default locale is EN and
// the reference time is fixed, so output never depends on the wall clock.
func New(seed uint64, opts ...Option) *Faker {
	c := config{locale: EN, ref: reference}
	for _, o := range opts {
		o(&c)
	}
	return &Faker{src: rand.NewPCG(seed, splitmix(seed)), loc: c.locale, ref: c.ref}
}

// Locale reports the active locale.
func (f *Faker) Locale() Locale { return f.loc }

// Clone returns an independent Faker seeded from this one's stream, sharing its
// locale and reference time. The sequence of clones is fixed by the parent's seed,
// so parallel work stays deterministic.
func (f *Faker) Clone() *Faker {
	return &Faker{src: rand.NewPCG(f.src.Uint64(), f.src.Uint64()), loc: f.loc, ref: f.ref}
}

// Fork returns n independent Fakers via Clone, for deterministic parallel seeding.
func (f *Faker) Fork(n int) []*Faker {
	out := make([]*Faker, n)
	for i := range out {
		out[i] = f.Clone()
	}
	return out
}

// splitmix derives a second seed word from one (SplitMix64 finalizer).
func splitmix(x uint64) uint64 {
	x += 0x9E3779B97F4A7C15
	x = (x ^ (x >> 30)) * 0xBF58476D1CE4E5B9
	x = (x ^ (x >> 27)) * 0x94D049BB133111EB
	return x ^ (x >> 31)
}
