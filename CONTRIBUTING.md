# Contributing

Issues and pull requests are welcome.

## Running the checks

```
go vet ./...
go test -race ./...
go test -coverpkg=./... -cover ./tests
```

All tests live in `tests/` and exercise the public API as a black box.

## The golden file

`tests/testdata/golden.txt` is the determinism contract: the output of a fixed seed
across the API. A change that alters it fails CI by design. If your change
*intentionally* alters generated output (which makes it a breaking change for a new
major version), regenerate it deliberately:

```
go test ./tests -run TestGolden -update
```

and call out the change in the pull request. If you did not mean to change output, a
golden failure means something shifted the stream; fix that instead of updating the
file.

## Adding a locale

Copy an existing locale subpackage, for example `locale/ng`, into `locale/<code>`:

```go
package xx

import "github.com/saphemmy/seedfaker"

var Locale = seedfaker.Locale{
	Code:    "en_XX",
	Country: "...",
	// names, phone formats, cities, states, streets, and the correct postcode shape
}

func init() { seedfaker.Register(Locale) }
```

Use real data, and get the postcode and phone formats right for the country. In a
format template, `#` is a digit and `?` is a letter. Add the package to the
blank-import list in `tests/locale_test.go` so it is checked.

## Style

`gofmt`, `go vet` and `golangci-lint` must pass. Keep documentation comments short
and in the standard form (a sentence that starts with the name). Add no dependencies.
