package tests

import (
	"testing"
	"time"

	"github.com/saphemmy/seedfaker"
	"github.com/saphemmy/seedfaker/locale/ng"
)

// TestEveryGeneratorProducesOutput calls the whole public surface once and checks
// that each result is non-empty and does not panic. It covers the generators the
// targeted tests do not reach.
func TestEveryGeneratorProducesOutput(t *testing.T) {
	f := seedfaker.New(123, seedfaker.WithLocale(ng.Locale))

	strs := map[string]func() string{
		"FirstName": f.FirstName, "LastName": f.LastName, "Name": f.Name,
		"Title": f.Title, "Suffix": f.Suffix, "FullName": f.FullName,
		"City": f.City, "State": f.State, "Street": f.Street, "Country": f.Country,
		"BuildingNumber": f.BuildingNumber, "Postcode": f.Postcode,
		"SecondaryAddress": f.SecondaryAddress, "StreetAddress": f.StreetAddress, "Address": f.Address,
		"Email": f.Email, "SafeEmail": f.SafeEmail, "Username": f.Username,
		"DomainWord": f.DomainWord, "TLD": f.TLD, "DomainName": f.DomainName, "URL": f.URL,
		"IPv4": f.IPv4, "IPv6": f.IPv6, "MACAddress": f.MACAddress, "Phone": f.Phone,
		"Company": f.Company, "CompanySuffix": f.CompanySuffix, "JobTitle": f.JobTitle,
		"CatchPhrase": f.CatchPhrase, "Bs": f.Bs,
		"CurrencyCode": f.CurrencyCode, "CreditCardType": f.CreditCardType, "CreditCardNumber": f.CreditCardNumber,
		"HexColor": f.HexColor, "RGBString": f.RGBString, "ColorName": f.ColorName,
		"EAN13": f.EAN13, "EAN8": f.EAN8, "ISBN10": f.ISBN10, "ISBN13": f.ISBN13,
		"Word": f.Word, "MD5": f.MD5, "SHA1": f.SHA1, "SHA256": f.SHA256, "Emoji": f.Emoji,
	}
	for name, gen := range strs {
		if gen() == "" {
			t.Errorf("%s returned empty", name)
		}
	}

	if f.Password(12) == "" || f.Slug(3) == "" || f.Words(4) == "" {
		t.Error("parameterised string generator returned empty")
	}
	if f.Sentences(2) == "" || f.Paragraph(2) == "" || f.Paragraphs(2) == "" || f.Text(100) == "" {
		t.Error("lorem generator returned empty")
	}

	// Numeric and time generators: just exercise them for panics and coverage.
	_ = f.Uint64()
	_ = f.Digit()
	_ = f.Letter()
	if m := f.Month(); m < time.January || m > time.December {
		t.Errorf("Month out of range: %v", m)
	}
	if wd := f.Weekday(); wd < time.Sunday || wd > time.Saturday {
		t.Errorf("Weekday out of range: %v", wd)
	}
	ref := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	if p := f.Past(24 * time.Hour); p.After(ref) {
		t.Error("Past returned a time after the reference")
	}
	if fut := f.Future(24 * time.Hour); fut.Before(ref) {
		t.Error("Future returned a time before the reference")
	}
	if !f.Between(ref, ref).Equal(ref) {
		t.Error("Between on a zero interval should return the bound")
	}
}

func TestWithReferenceTimeControlsDates(t *testing.T) {
	ref := time.Date(2000, time.June, 15, 0, 0, 0, 0, time.UTC)
	f := seedfaker.New(1, seedfaker.WithReferenceTime(ref))
	for range 500 {
		if d := f.Date(); d.After(ref) {
			t.Fatalf("Date %v is after the reference %v", d, ref)
		}
	}
}
