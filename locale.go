package seedfaker

import "sort"

// Locale holds the per-country data the locale-aware generators draw from. The
// built-in country locales live in their own subpackages (locale/ng, locale/gh,
// ...); import the ones you need and pass them with WithLocale, or build your own.
// A generator that needs a field the locale leaves empty panics with a clear
// message, so partial locales are fine as long as you only call what they support.
type Locale struct {
	Code            string // e.g. "en_NG"
	Country         string // e.g. "Nigeria"
	FirstNames      []string
	LastNames       []string
	Titles          []string // Mr, Mrs, Dr, ...
	PhoneFormats    []string // '#' -> digit, '?' -> A-Z letter, e.g. "+234 70# ### ####"
	Cities          []string
	States          []string // states, regions or provinces
	Streets         []string // street names, without the number
	PostcodeFormats []string // '#' -> digit, '?' -> letter, e.g. "######"
}

var registry = map[string]Locale{}

// Register adds a locale to the lookup used by ByCode. A locale subpackage calls it
// from init, so importing the package makes its locale resolvable by code.
// Registering a code twice panics, catching an accidental collision at startup.
func Register(l Locale) {
	if _, dup := registry[l.Code]; dup {
		panic("seedfaker: locale " + l.Code + " registered twice")
	}
	registry[l.Code] = l
}

// ByCode returns the locale registered for a code (e.g. "en_NG"), or ok=false. Only
// locales whose package has been imported are present.
func ByCode(code string) (l Locale, ok bool) { l, ok = registry[code]; return }

// Locales returns the codes of every registered locale, sorted.
func Locales() []string {
	out := make([]string, 0, len(registry))
	for c := range registry {
		out = append(out, c)
	}
	sort.Strings(out)
	return out
}

// expand fills a locale phone or postcode template: '#' -> digit, '?' -> letter.
func (f *Faker) expand(format string) string { return f.fill(format, true, true, false) }

func (f *Faker) require(field []string, name string) []string {
	if len(field) == 0 {
		panic("seedfaker: locale " + f.loc.Code + " has no " + name)
	}
	return field
}
