package seedfaker

import "strings"

// Numerify replaces each '#' in pattern with a digit 0-9. Other runes are copied.
func (f *Faker) Numerify(pattern string) string {
	return f.fill(pattern, true, false, false)
}

// Lexify replaces each '?' in pattern with an uppercase letter A-Z.
func (f *Faker) Lexify(pattern string) string {
	return f.fill(pattern, false, true, false)
}

// Bothify replaces '#' with a digit, '?' with a letter and '*' with either.
func (f *Faker) Bothify(pattern string) string {
	return f.fill(pattern, true, true, true)
}

func (f *Faker) fill(pattern string, hash, question, star bool) string {
	var b strings.Builder
	b.Grow(len(pattern))
	for _, r := range pattern {
		switch {
		case hash && r == '#':
			b.WriteByte(byte('0' + f.IntN(10)))
		case question && r == '?':
			b.WriteByte(byte('A' + f.IntN(26)))
		case star && r == '*':
			if f.Bool() {
				b.WriteByte(byte('0' + f.IntN(10)))
			} else {
				b.WriteByte(byte('A' + f.IntN(26)))
			}
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
