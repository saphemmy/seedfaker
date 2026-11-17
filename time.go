package seedfaker

import "time"

// Between returns a time uniformly in [start, end]; it panics if end < start.
func (f *Faker) Between(start, end time.Time) time.Time {
	d := end.Sub(start)
	if d < 0 {
		panic("seedfaker: Between requires start <= end")
	}
	if d == 0 {
		return start
	}
	return start.Add(time.Duration(f.u64n(uint64(d) + 1)))
}

// Past returns a time within d before the reference time.
func (f *Faker) Past(d time.Duration) time.Time { return f.Between(f.ref.Add(-d), f.ref) }

// Future returns a time within d after the reference time.
func (f *Faker) Future(d time.Duration) time.Time { return f.Between(f.ref, f.ref.Add(d)) }

// Date returns a date within the 30 years before the reference time.
func (f *Faker) Date() time.Time { return f.Past(30 * 365 * 24 * time.Hour) }

// Weekday returns a random weekday.
func (f *Faker) Weekday() time.Weekday { return time.Weekday(f.IntN(7)) }

// Month returns a random month.
func (f *Faker) Month() time.Month { return time.Month(f.IntN(12) + 1) }
