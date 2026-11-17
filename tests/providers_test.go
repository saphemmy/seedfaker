package tests

import (
	"strings"
	"testing"

	"github.com/saphemmy/seedfaker"
	"github.com/saphemmy/seedfaker/locale/ng"
)

func luhnValid(s string) bool {
	sum, double := 0, false
	for i := len(s) - 1; i >= 0; i-- {
		d := int(s[i] - '0')
		if double {
			if d *= 2; d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

// eanValid checks an EAN/ISBN-13 code: from the right, weights alternate 1,3,...
// so the rightmost (the check digit) has weight 1. This is length-independent,
// unlike weighting from the left.
func eanValid(s string) bool {
	sum, weight := 0, 1
	for i := len(s) - 1; i >= 0; i-- {
		sum += int(s[i]-'0') * weight
		weight = 4 - weight
	}
	return sum%10 == 0
}

func TestEmailIsSafeAndWellFormed(t *testing.T) {
	f := seedfaker.New(1)
	for range 1000 {
		e := f.Email()
		if !strings.Contains(e, "@") || e != strings.ToLower(e) {
			t.Fatalf("Email = %q", e)
		}
		if at := e[strings.LastIndex(e, "@")+1:]; !strings.HasPrefix(at, "example.") {
			t.Fatalf("Email %q not on a reserved domain", e)
		}
	}
}

func TestPhoneUsesLocalePrefix(t *testing.T) {
	f := seedfaker.New(1, seedfaker.WithLocale(ng.Locale))
	for range 500 {
		if p := f.Phone(); !strings.HasPrefix(p, "+234") {
			t.Fatalf("NG Phone = %q", p)
		}
	}
}

func TestCreditCardNumberPassesLuhn(t *testing.T) {
	f := seedfaker.New(2)
	for range 2000 {
		if n := f.CreditCardNumber(); len(n) != 16 || !luhnValid(n) {
			t.Fatalf("CreditCardNumber %q not Luhn-valid", n)
		}
	}
}

func TestBarcodesHaveValidCheckDigits(t *testing.T) {
	f := seedfaker.New(3)
	for range 1000 {
		if e := f.EAN13(); len(e) != 13 || !eanValid(e) {
			t.Fatalf("EAN13 %q invalid", e)
		}
		if e := f.EAN8(); len(e) != 8 || !eanValid(e) {
			t.Fatalf("EAN8 %q invalid", e)
		}
		if b := f.ISBN13(); len(b) != 13 || !strings.HasPrefix(b, "978") || !eanValid(b) {
			t.Fatalf("ISBN13 %q invalid", b)
		}
		if b := f.ISBN10(); len(b) != 10 {
			t.Fatalf("ISBN10 %q wrong length", b)
		}
	}
}

func TestUUIDIsVersion4AndCollisionFree(t *testing.T) {
	f := seedfaker.New(4)
	seen := make(map[string]struct{}, 20000)
	for range 20000 {
		u := f.UUID()
		if len(u) != 36 || u[14] != '4' {
			t.Fatalf("UUID %q not v4", u)
		}
		if u[8] != '-' || u[13] != '-' || u[18] != '-' || u[23] != '-' {
			t.Fatalf("UUID %q malformed", u)
		}
		if v := u[19]; v != '8' && v != '9' && v != 'a' && v != 'b' {
			t.Fatalf("UUID %q bad variant nibble %c", u, v)
		}
		if _, dup := seen[u]; dup {
			t.Fatalf("UUID collision at %q within 20k draws", u)
		}
		seen[u] = struct{}{}
	}
}

func TestHexColorAndRGB(t *testing.T) {
	f := seedfaker.New(5)
	for range 1000 {
		if c := f.HexColor(); len(c) != 7 || c[0] != '#' {
			t.Fatalf("HexColor = %q", c)
		}
		if r, g, b := f.RGB(); r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
			t.Fatalf("RGB = %d,%d,%d", r, g, b)
		}
	}
}

func TestLatLongRanges(t *testing.T) {
	f := seedfaker.New(6)
	for range 2000 {
		if lat := f.Latitude(); lat < -90 || lat >= 90 {
			t.Fatalf("Latitude = %v", lat)
		}
		if lng := f.Longitude(); lng < -180 || lng >= 180 {
			t.Fatalf("Longitude = %v", lng)
		}
	}
}

func TestIPv6AndMACAreHexGroups(t *testing.T) {
	f := seedfaker.New(7)
	hex := func(s string) bool {
		for _, r := range s {
			isHex := (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f')
			if !isHex {
				return false
			}
		}
		return true
	}
	for range 500 {
		if g := strings.Split(f.IPv6(), ":"); len(g) != 8 || !hex(strings.Join(g, "")) {
			t.Fatalf("IPv6 = %q", f.IPv6())
		}
		if g := strings.Split(f.MACAddress(), ":"); len(g) != 6 {
			t.Fatalf("MAC groups = %d", len(g))
		}
	}
}

func TestTextEdges(t *testing.T) {
	f := seedfaker.New(8)
	if w := f.Words(0); w != "" {
		t.Fatalf("Words(0) = %q", w)
	}
	if s := f.Sentence(0); !strings.HasSuffix(s, ".") {
		t.Fatalf("Sentence(0) = %q", s)
	}
	if txt := f.Text(200); len(txt) > 200 {
		t.Fatalf("Text(200) too long: %d", len(txt))
	}
	// A maxChars shorter than any word still yields exactly one (whole) word.
	if txt := f.Text(2); txt == "" || strings.Contains(txt, " ") {
		t.Fatalf("Text(2) should be one whole word, got %q", txt)
	}
	assertPanics(t, "Text(0)", func() { f.Text(0) })
	assertPanics(t, "Password(0)", func() { f.Password(0) })
	assertPanics(t, "Boolean(101)", func() { f.Boolean(101) })
	assertPanics(t, "Boolean(-1)", func() { f.Boolean(-1) })
}

func TestProviderPanicsOnEmptyLocaleField(t *testing.T) {
	f := seedfaker.New(1, seedfaker.WithLocale(seedfaker.Locale{Code: "xx"}))
	assertPanics(t, "City", func() { f.City() })
	assertPanics(t, "Phone", func() { f.Phone() })
	assertPanics(t, "Country", func() { f.Country() })
	assertPanics(t, "Postcode", func() { f.Postcode() })
}
