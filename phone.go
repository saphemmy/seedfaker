package seedfaker

// Phone returns a number in one of the locale's formats; it panics if the locale
// defines none.
func (f *Faker) Phone() string {
	if len(f.loc.PhoneFormats) == 0 {
		panic("seedfaker: locale " + f.loc.Code + " has no phone formats")
	}
	return f.expand(f.pick(f.loc.PhoneFormats))
}
