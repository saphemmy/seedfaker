package tests

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/saphemmy/seedfaker"
	"github.com/saphemmy/seedfaker/locale/ng"
)

var update = flag.Bool("update", false, "rewrite the golden file")

// snapshot exercises the whole API from a fixed seed and locale. Its output is the
// determinism contract: it must not change except on a deliberate major release.
func snapshot() string {
	f := seedfaker.New(42, seedfaker.WithLocale(ng.Locale))
	var b strings.Builder
	w := func(label, v string) { b.WriteString(label + "\t" + v + "\n") }
	w("Name", f.Name())
	w("FullName", f.FullName())
	w("Title", f.Title())
	w("Email", f.Email())
	w("Username", f.Username())
	w("Phone", f.Phone())
	w("UUID", f.UUID())
	w("StreetAddress", f.StreetAddress())
	w("City", f.City())
	w("State", f.State())
	w("Postcode", f.Postcode())
	w("Country", f.Country())
	w("Company", f.Company())
	w("JobTitle", f.JobTitle())
	w("CatchPhrase", f.CatchPhrase())
	w("CreditCardType", f.CreditCardType())
	w("CreditCardNumber", f.CreditCardNumber())
	w("CurrencyCode", f.CurrencyCode())
	w("HexColor", f.HexColor())
	w("ColorName", f.ColorName())
	w("EAN13", f.EAN13())
	w("ISBN10", f.ISBN10())
	w("ISBN13", f.ISBN13())
	w("Word", f.Word())
	w("Sentence", f.Sentence(8))
	w("IPv4", f.IPv4())
	w("IPv6", f.IPv6())
	w("MACAddress", f.MACAddress())
	w("MD5", f.MD5())
	w("Bothify", f.Bothify("ID-####-???"))
	w("Date", f.Date().UTC().Format(time.RFC3339))
	return b.String()
}

// TestGolden fails if the output of a fixed seed drifts. Regenerate deliberately
// with: go test ./tests -run TestGolden -update
func TestGolden(t *testing.T) {
	got := snapshot()
	path := filepath.Join("testdata", "golden.txt")
	if *update {
		if err := os.MkdirAll("testdata", 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}
	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read golden (first run: go test ./tests -run TestGolden -update): %v", err)
	}
	if got != string(want) {
		t.Errorf("output drifted from the golden file; if intended this is a major release.\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

func TestSameSeedSameOutput(t *testing.T) {
	first, second := snapshot(), snapshot()
	if first != second {
		t.Fatal("same seed produced different output")
	}
}

func TestDifferentSeedsDiffer(t *testing.T) {
	var a, b strings.Builder
	fa, fb := seedfaker.New(1), seedfaker.New(2)
	for range 20 {
		a.WriteString(fa.Name())
		b.WriteString(fb.Name())
	}
	if a.String() == b.String() {
		t.Fatal("different seeds produced identical streams")
	}
}

func TestCloneReproducesFromEqualSeeds(t *testing.T) {
	a := seedfaker.New(7).Clone()
	b := seedfaker.New(7).Clone()
	for range 50 {
		if a.IntN(1000) != b.IntN(1000) {
			t.Fatal("clones from equally seeded parents diverged")
		}
	}
}

func TestForkIsDeterministic(t *testing.T) {
	a := seedfaker.New(11).Fork(3)
	b := seedfaker.New(11).Fork(3)
	for i := range a {
		if a[i].UUID() != b[i].UUID() {
			t.Fatalf("fork %d diverged across equal seeds", i)
		}
	}
}

func TestDateNeverReadsTheWallClock(t *testing.T) {
	d1, d2 := seedfaker.New(99).Date(), seedfaker.New(99).Date()
	if !d1.Equal(d2) {
		t.Fatal("Date is not deterministic")
	}
}
