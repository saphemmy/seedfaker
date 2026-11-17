package seedfaker

import "time"

type config struct {
	locale Locale
	ref    time.Time
}

// Option configures a Faker.
type Option func(*config)

// WithLocale sets the locale; the default is EN.
func WithLocale(l Locale) Option { return func(c *config) { c.locale = l } }

// WithReferenceTime sets the "now" that Date, Past and Future measure from, keeping
// output deterministic regardless of the wall clock. The default is fixed.
func WithReferenceTime(t time.Time) Option { return func(c *config) { c.ref = t } }
