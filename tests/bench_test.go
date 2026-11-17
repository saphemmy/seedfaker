package tests

import (
	"testing"

	"github.com/saphemmy/seedfaker"
)

func BenchmarkIntN(b *testing.B) {
	f := seedfaker.New(1)
	b.ReportAllocs()
	for range b.N {
		_ = f.IntN(1000)
	}
}

func BenchmarkName(b *testing.B) {
	f := seedfaker.New(1)
	b.ReportAllocs()
	for range b.N {
		_ = f.Name()
	}
}

func BenchmarkUUID(b *testing.B) {
	f := seedfaker.New(1)
	b.ReportAllocs()
	for range b.N {
		_ = f.UUID()
	}
}

func BenchmarkFill(b *testing.B) {
	f := seedfaker.New(1)
	b.ReportAllocs()
	for range b.N {
		var u user
		_ = f.Fill(&u)
	}
}
