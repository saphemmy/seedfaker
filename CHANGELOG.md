# Changelog

The format follows [Keep a Changelog](https://keepachangelog.com/), and the project
follows [Semantic Versioning](https://semver.org/).

## [0.1.0] - 2026-11-17

Initial release.

### Added
- Deterministic `Faker`: a PCG source with an in-library uniform reduction, a fixed
  reference time, and no global state — one seed yields one dataset.
- Primitives: `IntN`, `IntBetween`, `Uint64`, `Bool`, `Float64`, `FloatBetween`,
  `Number`, `DigitNot`, `DigitNotNull`, `Letter`, `Pick`, `Shuffle`, `Weighted`,
  and the `Numerify`/`Lexify`/`Bothify` template fillers.
- Bounded, error-returning modifiers: `UniqueN`, `Optional`, `Valid`.
- Providers: person, address, phone, company, internet, finance (Luhn-valid cards
  and currencies), colour, barcode (valid EAN/ISBN check digits), date and time,
  lorem, UUID, hashes and miscellaneous.
- 13 country locales as self-registering subpackages, plus the `EN` default.
- `Factory[T]` and `Fill` for ORM-agnostic seeding.
- A golden-file determinism test, diffed in CI across a Go-version matrix.
