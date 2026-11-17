package seedfaker

import "fmt"

var secondaryAddressFormats = []string{"Apt. ###", "Suite ###", "Unit ###", "Floor ##"}

// City returns a city for the active locale.
func (f *Faker) City() string { return f.pick(f.require(f.loc.Cities, "cities")) }

// State returns a state, region or province for the active locale.
func (f *Faker) State() string { return f.pick(f.require(f.loc.States, "states")) }

// Street returns a street name for the active locale.
func (f *Faker) Street() string { return f.pick(f.require(f.loc.Streets, "streets")) }

// Country returns the active locale's country name.
func (f *Faker) Country() string {
	if f.loc.Country == "" {
		panic("seedfaker: locale " + f.loc.Code + " has no country")
	}
	return f.loc.Country
}

// BuildingNumber returns a 1-4 digit building number.
func (f *Faker) BuildingNumber() string { return fmt.Sprintf("%d", f.IntBetween(1, 9999)) }

// Postcode returns a postcode in one of the locale's formats.
func (f *Faker) Postcode() string {
	return f.expand(f.pick(f.require(f.loc.PostcodeFormats, "postcode formats")))
}

// SecondaryAddress returns a unit line such as "Apt. 123".
func (f *Faker) SecondaryAddress() string { return f.expand(f.pick(secondaryAddressFormats)) }

// StreetAddress returns a building number and street, e.g. "123 Main Street".
func (f *Faker) StreetAddress() string { return f.BuildingNumber() + " " + f.Street() }

// Address returns a full address: street, city, state and postcode.
func (f *Faker) Address() string {
	return fmt.Sprintf("%s, %s, %s %s", f.StreetAddress(), f.City(), f.State(), f.Postcode())
}

// Latitude returns a latitude in [-90, 90).
func (f *Faker) Latitude() float64 { return f.FloatBetween(-90, 90) }

// Longitude returns a longitude in [-180, 180).
func (f *Faker) Longitude() float64 { return f.FloatBetween(-180, 180) }
