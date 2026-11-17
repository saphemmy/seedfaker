# seedfaker

[![CI](https://github.com/saphemmy/seedfaker/actions/workflows/ci.yml/badge.svg)](https://github.com/saphemmy/seedfaker/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/saphemmy/seedfaker.svg)](https://pkg.go.dev/github.com/saphemmy/seedfaker)
[![Go Report Card](https://goreportcard.com/badge/github.com/saphemmy/seedfaker)](https://goreportcard.com/report/github.com/saphemmy/seedfaker)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Deterministic fake data for Go. The same seed produces the same data on every run, every machine and every supported Go version — so seeded fixtures, golden tests and re-runnable database seeders stay stable instead of drifting.

```go
go get github.com/saphemmy/seedfaker
```

Go 1.22 or later. No dependencies.

## Why this exists

seedfaker came out of building and re-seeding real services. I kept needing test data that looked real and, more importantly, stayed the same every time a seeder ran — otherwise an idempotent seeder isn't idempotent and a golden test isn't golden. I also needed data that fit the people the software actually served: Nigerian and Ghanaian names, phone numbers and postcodes, which the existing Go fakers get wrong or simply don't have.

None of those fakers would promise that the same seed gives the same bytes. They seed a package-global RNG, draw UUIDs from `crypto/rand`, and read the wall clock inside their date helpers — so a "seed" fixes some values and quietly not others. Seeded data that doesn't reproduce is worse than no seed at all: it looks repeatable, then changes under you. So I wrote one that treats reproducibility as the product, ships the locales I needed first, and carries no dependencies.

| Property | What it buys you |
|---|---|
| **One seed, one dataset** | A `Faker` draws every value — including UUIDs and dates — from a single seeded source. Same seed in, same bytes out. |
| **No wall clock** | Date generators measure from a fixed reference time, never `time.Now()`, so they reproduce. |
| **Stable across Go versions** | Values come from a PCG source plus an in-library reduction, not `math/rand` internals that change between releases. A golden file in CI proves it on every supported Go version. |
| **No global state** | No package-level RNG, no hidden singletons. Two fakers never interfere. |
| **Zero dependencies** | Nothing in your dependency tree but the standard library. |
| **Deterministic parallelism** | `Clone` and `Fork` hand out independent fakers whose streams are fixed by the parent's seed. |

What it is **not**: a cryptographic source. seedfaker is predictable by design — never use it for passwords, tokens or anything that must be unguessable.

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

	fmt.Println(f.Name())   // Nigerian name, same for seed 42 every run
	fmt.Println(f.Email())  // always on a reserved domain — never a real inbox
	fmt.Println(f.Phone())  // +234 ...
	fmt.Println(f.UUID())   // deterministic v4 UUID
}
```

## The determinism contract

Within a major version, a given seed produces identical output forever. Richer data ships as *new* locales or generators, never by changing existing data (which would shift the stream). `tests/testdata/golden.txt` captures the output of a fixed seed across the whole API, and CI diffs it on every supported Go version; any drift fails the build, so the guarantee is enforced mechanically rather than promised. See [COMPATIBILITY.md](COMPATIBILITY.md).

## Locales

Country data lives in small subpackages, so you compile in only the locales you use:

```go
import "github.com/saphemmy/seedfaker/locale/gh" // Ghana

f := faker.New(1, faker.WithLocale(gh.Locale))
```

Built in: `ng` `gh` `ke` `za` `us` `gb` `fr` `es` `de` `br` `in` `jp` `cn`, with the locale-neutral `EN` as the default. Each carries names, phone formats, cities, states, streets and the correct postcode shape for its country (UK letter-postcodes, Ghana GhanaPost codes, US ZIP+4, and so on). Build your own `Locale` for anywhere else, and register it with `faker.Register` to resolve it by code.

## Generators

Person, address, phone, company, internet, finance (Luhn-valid cards, currencies), colour, barcode (valid EAN/ISBN check digits), date and time, lorem, UUID, hashes and more — plus the primitives (`IntN`, `Float64`, `Pick`, `Shuffle`, `Weighted`, `Numerify`/`Lexify`/`Bothify`) that every generator is built on, so a custom one is exactly as deterministic as the built-ins. The modifiers `UniqueN`, `Optional` and `Valid` are bounded and return an error on exhaustion instead of looping.

## Factories and your database

Factories are the ORM-agnostic way to seed. You describe how to build one record; `Create` persists each through a callback, so the library stays free of any ORM:

```go
users := faker.Factory[User]{
	Build: func(f *faker.Faker) User {
		return User{Name: f.Name(), Email: f.Email()}
	},
}

// Make in memory, or Create and persist through GORM, ent, bun, sqlx, raw SQL — anything.
list := users.Make(faker.New(1), 100)
_, err := users.Create(faker.New(1), 100, func(u User) error {
	return db.Create(&u).Error
})
```

`Fill` populates an existing struct — including an ORM model — from `faker:"..."` struct tags.

## Testing

```
go test ./...
go test -race ./...
go test -coverpkg=./... -cover ./tests
```

Every test lives in `tests/`. The suite covers the public API as a black box, including the golden determinism file, property and edge-case tests, and fuzz targets for the primitives.

## Versioning

Pre-1.0. The API may gain additive changes on minor versions; generated output for a seed is frozen within a major version. The first stable release will lock both.

## Contributing

Issues and pull requests are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md). New locales are especially welcome and are easy to add — one small subpackage.

## License

MIT © Oluwafemi Sosami
