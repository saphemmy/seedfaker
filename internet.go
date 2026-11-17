package seedfaker

import (
	"fmt"
	"strings"
)

var (
	// reservedDomains (RFC 2606) never resolve, so a generated email can never
	// reach a real inbox — a safety guarantee fakerphp does not make.
	reservedDomains = []string{"example.com", "example.net", "example.org", "example.test"}
	domainWords     = []string{"acme", "globex", "initech", "umbrella", "hooli", "example", "demo", "sample"}
	tlds            = []string{"com", "net", "org", "io", "co", "dev", "app"}
	passwordChars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Email returns an address on a reserved domain, derived from a locale name, so it
// can never reach a real inbox.
func (f *Faker) Email() string {
	user := strings.ToLower(f.FirstName() + "." + f.LastName())
	return fmt.Sprintf("%s.%d@%s", user, f.IntN(10000), f.pick(reservedDomains))
}

// SafeEmail is an alias for Email; every generated address is already safe.
func (f *Faker) SafeEmail() string { return f.Email() }

// Username returns a lowercase handle derived from a locale name.
func (f *Faker) Username() string {
	return fmt.Sprintf("%s%d", strings.ToLower(f.FirstName()), f.IntN(1000))
}

// Password returns a random alphanumeric password of the given length; it panics if
// length < 1.
func (f *Faker) Password(length int) string {
	if length < 1 {
		panic("seedfaker: Password requires length >= 1")
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = passwordChars[f.IntN(len(passwordChars))]
	}
	return string(b)
}

// DomainWord returns a single lowercase domain label.
func (f *Faker) DomainWord() string { return f.pick(domainWords) }

// TLD returns a top-level domain without the dot.
func (f *Faker) TLD() string { return f.pick(tlds) }

// DomainName returns a synthetic domain, e.g. "acme.io".
func (f *Faker) DomainName() string { return f.DomainWord() + "." + f.TLD() }

// URL returns a synthetic https URL.
func (f *Faker) URL() string { return "https://www." + f.DomainName() + "/" + f.Slug(3) }

// Slug returns a hyphenated lowercase slug of n lorem words (n <= 0 gives 3).
func (f *Faker) Slug(n int) string {
	if n <= 0 {
		n = 3
	}
	p := make([]string, n)
	for i := range p {
		p[i] = f.Word()
	}
	return strings.Join(p, "-")
}

// IPv4 returns a dotted IPv4 address.
func (f *Faker) IPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", f.IntN(256), f.IntN(256), f.IntN(256), f.IntN(256))
}

// IPv6 returns a colon-separated IPv6 address.
func (f *Faker) IPv6() string {
	g := make([]string, 8)
	for i := range g {
		g[i] = fmt.Sprintf("%04x", f.IntN(1<<16))
	}
	return strings.Join(g, ":")
}

// MACAddress returns a colon-separated MAC address.
func (f *Faker) MACAddress() string {
	g := make([]string, 6)
	for i := range g {
		g[i] = fmt.Sprintf("%02x", f.IntN(256))
	}
	return strings.Join(g, ":")
}
