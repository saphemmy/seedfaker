# Compatibility and the determinism contract

seedfaker's value is that a seed reproduces. This document states exactly what is
guaranteed, so you can depend on it.

## The guarantee

Within a major version, a given seed and locale produce identical output on every
run, every machine and every supported Go version.

Three design choices make this true:

1. **One source.** Every value (numbers, strings, UUIDs, dates) comes from a single
   PCG generator seeded from your seed. There is no second, unseeded source.
2. **Our own reduction.** Bounded integers come from a reduction implemented in this
   library (Lemire's method over the raw PCG output), not from `math/rand`, whose
   distribution helpers may change between Go releases.
3. **A fixed reference time.** Date generators measure from a reference time that
   defaults to a fixed instant and is otherwise set explicitly. Nothing reads the
   wall clock.

## How it is enforced

`tests/testdata/golden.txt` holds the output of a fixed seed across the whole API.
The test suite regenerates that output and diffs it, and CI runs the suite on every
supported Go version. Any change to generated output fails the build. Changing the
golden file is therefore a deliberate act, reserved for a major release.

## What may change, and when

- **Patch and minor releases** may add locales, generators and helpers. They do not
  change the output of existing generators for an existing seed.
- **Adding data to an existing locale** (for example a name) shifts that locale's
  stream, so it is a breaking change and ships only in a major release. New data for
  an existing country goes into the release notes.
- **Major releases** may change output. They are the only releases that do.

## Supported Go versions

The two most recent Go releases, plus the minimum stated in `go.mod` (currently
1.22). The CI matrix runs the golden test on each.

## Scope

seedfaker is not a cryptographic generator. Its output is predictable by design.
