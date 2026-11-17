package seedfaker

import "strings"

var (
	currencyCodes   = []string{"USD", "EUR", "GBP", "NGN", "GHS", "JPY", "CNY", "INR", "ZAR", "KES", "BRL", "CAD", "AUD"}
	creditCardTypes = []string{"Visa", "MasterCard", "American Express", "Discover"}
)

// CurrencyCode returns an ISO 4217 currency code.
func (f *Faker) CurrencyCode() string { return f.pick(currencyCodes) }

// CreditCardType returns a card scheme name.
func (f *Faker) CreditCardType() string { return f.pick(creditCardTypes) }

// CreditCardNumber returns a 16-digit, Luhn-valid Visa test number (prefix 4). It
// is synthetic: valid in form, never a real account.
func (f *Faker) CreditCardNumber() string {
	d := make([]int, 16)
	d[0] = 4
	for i := 1; i < 15; i++ {
		d[i] = f.IntN(10)
	}
	d[15] = luhnCheckDigit(d[:15])
	var b strings.Builder
	for _, x := range d {
		b.WriteByte(byte('0' + x))
	}
	return b.String()
}

// luhnCheckDigit returns the digit that makes body+digit pass the Luhn check.
func luhnCheckDigit(body []int) int {
	sum, double := 0, true
	for i := len(body) - 1; i >= 0; i-- {
		x := body[i]
		if double {
			if x *= 2; x > 9 {
				x -= 9
			}
		}
		sum += x
		double = !double
	}
	return (10 - sum%10) % 10
}
