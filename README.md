# seedfaker

[![CI](https://github.com/saphemmy/seedfaker/actions/workflows/ci.yml/badge.svg)](https://github.com/saphemmy/seedfaker/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/saphemmy/seedfaker.svg)](https://pkg.go.dev/github.com/saphemmy/seedfaker)
[![Go Report Card](https://goreportcard.com/badge/github.com/saphemmy/seedfaker)](https://goreportcard.com/report/github.com/saphemmy/seedfaker)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Deterministic fake data for Go. The same seed produces the same data on every run, machine, and supported Go version, so seeded fixtures, golden tests, and re-runnable database seeders stay stable.

```go
go get github.com/saphemmy/seedfaker
```

Go 1.22 or later. No dependencies.

## Why this exists

I wrote seedfaker while building database seeders for real services. A seeder that runs more than once has to produce the same records each time, or it creates duplicates and the data drifts. That needs a faker whose output follows entirely from a seed.

The fakers I tried did not work that way. They seed a package-global random source, generate UUIDs from `crypto/rand`, and read the system clock in their date helpers, so the same seed fixes some values and leaves others free. I also needed Nigerian and Ghanaian data, such as names, phone numbers, and postcodes, that they got wrong or did not have.

seedfaker is built around one rule: a seed fully determines the output. Every value comes from a single seeded source, dates come from a reference time instead of the clock, and there is no global state. It has no dependencies, and the locales I needed ship with it.

## Design

| Property | Detail |
|---|---|
| One seeded source | Every value, including UUIDs and dates, is drawn from one PCG generator seeded from your seed. |
| No wall clock | Date generators measure from a reference time that defaults to a fixed instant and can be set with `WithReferenceTime`. |
| Stable across Go versions | Bounded values use a reduction implemented in this package, not `math/rand` helpers that can change between releases. A golden file in CI checks this on each supported version. |
| No global state | No package-level RNG and no singletons. Two fakers never interfere. |
| No dependencies | Standard library only. |
| Deterministic parallelism | `Clone` and `Fork` return independent fakers whose streams follow from the parent's seed. |

seedfaker is not a cryptographic source. Its output is predictable by design, so do not use it for passwords, tokens, or anything that must be unguessable.

## Quick start

```go
package main

import (
	"fmt"

	faker "github.com/saphemmy/seedfaker"
	"github.com/saphemmy/seedfaker/locale/ng"
)

func main() {
	f := faker.New(42, faker.WithLocale(ng.Locale))

	fmt.Println(f.Name())   // Nigerian name, same for seed 42 on every run
	fmt.Println(f.Email())  // always on a reserved domain, never a real inbox
	fmt.Println(f.Phone())  // +234 ...
	fmt.Println(f.UUID())   // deterministic v4 UUID
}
```

## Determinism

Within a major version, a given seed and locale produce the same output. New data ships as new locales or generators rather than changes to existing data, which would shift the stream. `tests/testdata/golden.txt` holds the output of a fixed seed across the whole API, and CI diffs it on each supported Go version, so a change to generated output fails the build. See [COMPATIBILITY.md](COMPATIBILITY.md).

## Locales

Country data lives in small subpackages, so a program compiles in only the locales it uses:

```go
import "github.com/saphemmy/seedfaker/locale/gh" // Ghana

f := faker.New(1, faker.WithLocale(gh.Locale))
```

Built in: `ng` `gh` `ke` `za` `us` `gb` `fr` `es` `de` `br` `in` `jp` `cn`, with the locale-neutral `EN` as the default. Each carries names, phone formats, cities, states, streets, and the correct postcode shape for its country, including UK letter-postcodes, GhanaPost codes, and US ZIP+4. Build your own `Locale` for anywhere else and call `faker.Register` to resolve it by code.

## Generators

Person, address, phone, company, internet, finance (Luhn-valid cards and currencies), colour, barcode (valid EAN and ISBN check digits), date and time, lorem, UUID, and hashes. Each is built on the exported primitives (`IntN`, `Float64`, `Pick`, `Shuffle`, `Weighted`, and the `Numerify`/`Lexify`/`Bothify` fillers), so a generator you write yourself is as deterministic as the built-in ones. The modifiers `UniqueN`, `Optional`, and `Valid` are bounded and return an error when the value space runs out rather than looping.

## Factories and your database

A factory describes how to build one record. `Create` persists each one through a callback you provide, so the library depends on no ORM:

```go
users := faker.Factory[User]{
	Build: func(f *faker.Faker) User {
		return User{Name: f.Name(), Email: f.Email()}
	},
}

list := users.Make(faker.New(1), 100) // in memory
_, err := users.Create(faker.New(1), 100, func(u User) error {
	return db.Create(&u).Error // or ent, bun, sqlx, raw SQL
})
```

`Fill` populates an existing struct, including an ORM model, from `faker:"..."` struct tags.

## Testing

```
go test ./...
go test -race ./...
go test -coverpkg=./... -cover ./tests
```

Every test lives in `tests/` and exercises the public API as a black box, including the golden determinism file, edge-case tests, and fuzz targets for the primitives.

## Versioning

Pre-1.0. Minor versions may add to the API. Generated output for a seed is fixed within a major version. The first stable release will fix both.

## Contributing

Issues and pull requests are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md). New locales are easy to add and especially welcome.

## License

MIT © Oluwafemi Sosami
