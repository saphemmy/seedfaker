package seedfaker

var nameSuffixes = []string{"Jr.", "Sr.", "II", "III", "IV", "PhD", "MD", "DDS"}

// FirstName returns a first name for the active locale.
func (f *Faker) FirstName() string { return f.pick(f.require(f.loc.FirstNames, "first names")) }

// LastName returns a surname for the active locale.
func (f *Faker) LastName() string { return f.pick(f.require(f.loc.LastNames, "last names")) }

// Name returns "First Last" for the active locale.
func (f *Faker) Name() string { return f.FirstName() + " " + f.LastName() }

// Title returns a personal title (Mr, Mrs, Dr, ...) for the active locale.
func (f *Faker) Title() string { return f.pick(f.require(f.loc.Titles, "titles")) }

// Suffix returns a name suffix such as "Jr." or "PhD".
func (f *Faker) Suffix() string { return f.pick(nameSuffixes) }

// FullName returns a title, first name and last name, e.g. "Dr Jane Smith".
func (f *Faker) FullName() string { return f.Title() + " " + f.FirstName() + " " + f.LastName() }
