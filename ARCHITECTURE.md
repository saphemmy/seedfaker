# Architecture

seedfaker is a small deterministic engine with everything else composed on top. The
core never grows a dependency, and first-party and third-party extensions sit at the
same level.

## The engine

`Faker` holds three things: a PCG source (`math/rand/v2`), the active `Locale`, and a
reference time. `New(seed)` seeds the PCG from the seed and a SplitMix64-derived
second word, so a caller passes one number.

Every bounded draw goes through `IntN`, which reduces the raw 64-bit PCG output with
Lemire's multiply-shift method, implemented here. Nothing calls `math/rand`'s
distribution helpers, whose output may change between Go releases. That one decision
is what makes output stable across Go versions; the PCG algorithm itself is fixed.

Dates measure from the reference time, never `time.Now()`. `Clone` seeds a child from
the parent's stream, so parallel work is deterministic in the parent's seed.

## Layers

- **Primitives** (`random.go`, `numbers.go`, `pick.go`, `format.go`): `IntN`,
  `Float64`, `Pick`, `Shuffle`, `Weighted`, the template fillers. Exported, so a
  custom generator is as deterministic as a built-in.
- **Providers** (`person.go`, `address.go`, `internet.go`, ...): methods on `Faker`.
  They are code, not data, so they live in the root package and add no weight.
- **Locales** (`locale/<code>/`): data, where weight actually lives. Each is a small
  self-registering subpackage, so a program compiles in only the countries it uses.
- **Composition** (`factory.go`, `fill.go`, `unique.go`): `Factory[T]`, `Fill`, and
  the bounded modifiers.

## Extending

- A **custom generator** is a plain `func(*Faker) T` built on the exported
  primitives. It needs no hook into the type.
- A **custom locale** is a `Locale` value; call `Register` to resolve it by code.
- A **named generator** for `Fill`'s `faker:"..."` tags is added with
  `RegisterGenerator`.
- An **ORM** is reached through `Factory.Create`'s persist callback; the core depends
  on none.

## Why locales are subpackages but providers are not

Providers are a few lines of code each; subpackaging them would buy nothing and cost
ergonomics. Locale data is where size grows with coverage, and most programs want one
or two countries, so that is what gets split — the same reasoning as
`golang.org/x/text`.
