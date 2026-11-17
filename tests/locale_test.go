package tests

import (
	"sort"
	"strings"
	"testing"

	"github.com/saphemmy/seedfaker"

	// Register every built-in locale for the catalogue tests.
	_ "github.com/saphemmy/seedfaker/locale/br"
	_ "github.com/saphemmy/seedfaker/locale/cn"
	_ "github.com/saphemmy/seedfaker/locale/de"
	_ "github.com/saphemmy/seedfaker/locale/es"
	_ "github.com/saphemmy/seedfaker/locale/fr"
	_ "github.com/saphemmy/seedfaker/locale/gb"
	_ "github.com/saphemmy/seedfaker/locale/gh"
	_ "github.com/saphemmy/seedfaker/locale/in"
	_ "github.com/saphemmy/seedfaker/locale/jp"
	_ "github.com/saphemmy/seedfaker/locale/ke"
	_ "github.com/saphemmy/seedfaker/locale/ng"
	_ "github.com/saphemmy/seedfaker/locale/us"
	_ "github.com/saphemmy/seedfaker/locale/za"
)

func TestLocalesAreRegisteredAndSorted(t *testing.T) {
	codes := seedfaker.Locales()
	if len(codes) < 14 { // EN default + 13 countries
		t.Fatalf("expected >=14 locales, got %d: %v", len(codes), codes)
	}
	if !sort.StringsAreSorted(codes) {
		t.Fatalf("Locales() is not sorted: %v", codes)
	}
}

func TestByCode(t *testing.T) {
	if _, ok := seedfaker.ByCode("en_NG"); !ok {
		t.Fatal("en_NG not resolvable")
	}
	if _, ok := seedfaker.ByCode("nonsense"); ok {
		t.Fatal("unknown code resolved")
	}
}

func TestEveryLocaleGeneratesValidPhoneAndPostcode(t *testing.T) {
	for _, code := range seedfaker.Locales() {
		loc, _ := seedfaker.ByCode(code)
		f := seedfaker.New(1, seedfaker.WithLocale(loc))
		for range 200 {
			if p := f.Phone(); !strings.HasPrefix(p, "+") {
				t.Fatalf("%s Phone = %q", code, p)
			}
			pc := f.Postcode()
			if pc == "" || strings.ContainsAny(pc, "#?") {
				t.Fatalf("%s Postcode = %q (template not expanded)", code, pc)
			}
		}
	}
}

func TestUKPostcodeHasLettersAndDigits(t *testing.T) {
	gb, _ := seedfaker.ByCode("en_GB")
	f := seedfaker.New(1, seedfaker.WithLocale(gb))
	letters, digits := false, false
	for range 100 {
		for _, r := range f.Postcode() {
			switch {
			case r >= 'A' && r <= 'Z':
				letters = true
			case r >= '0' && r <= '9':
				digits = true
			}
		}
	}
	if !letters || !digits {
		t.Fatal("UK postcodes should contain both letters and digits")
	}
}

func TestRegisterRejectsDuplicateCode(t *testing.T) {
	ng, _ := seedfaker.ByCode("en_NG")
	assertPanics(t, "duplicate locale", func() { seedfaker.Register(ng) })
}
