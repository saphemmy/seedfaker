package seedfaker

import (
	"encoding/binary"
	"encoding/hex"
)

// UUID returns a deterministic RFC 4122 version-4 UUID. Fakers that draw UUIDs
// from crypto/rand cannot reproduce; this one is seeded, so it does.
func (f *Faker) UUID() string {
	var b [16]byte
	binary.BigEndian.PutUint64(b[0:8], f.src.Uint64())
	binary.BigEndian.PutUint64(b[8:16], f.src.Uint64())
	b[6] = b[6]&0x0f | 0x40 // version 4
	b[8] = b[8]&0x3f | 0x80 // variant RFC 4122
	var dst [36]byte
	hex.Encode(dst[0:8], b[0:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], b[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], b[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], b[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:36], b[10:16])
	return string(dst[:])
}
