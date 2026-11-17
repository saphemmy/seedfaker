package seedfaker

import "encoding/hex"

var emojis = []string{
	"\U0001F600", "\U0001F602", "\U0001F60D", "\U0001F44D", "\U0001F389", "\U0001F680",
	"\U0001F525", "✨", "\U0001F4A1", "\U0001F4E6", "\U0001F30D", "⭐",
}

// Boolean reports true with the given percent chance (0-100); it panics if percent
// is out of range.
func (f *Faker) Boolean(percent int) bool {
	if percent < 0 || percent > 100 {
		panic("seedfaker: Boolean requires percent in 0..100")
	}
	return f.IntN(100) < percent
}

// hexOf returns n random bytes hex-encoded (2n characters).
func (f *Faker) hexOf(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(f.IntN(256))
	}
	return hex.EncodeToString(b)
}

// MD5 returns a random 32-character hex string (the shape of an md5 digest).
func (f *Faker) MD5() string { return f.hexOf(16) }

// SHA1 returns a random 40-character hex string.
func (f *Faker) SHA1() string { return f.hexOf(20) }

// SHA256 returns a random 64-character hex string.
func (f *Faker) SHA256() string { return f.hexOf(32) }

// Emoji returns a single emoji.
func (f *Faker) Emoji() string { return f.pick(emojis) }
