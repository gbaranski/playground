package scanstring

import (
	"math"
	"math/rand"
	"testing"
)

var (
	input       []byte
	occurencies uint
)

// 5MB
const inputLength = 1000 << (10 * 2)
const wantedByte = 10

func init() {

	input = make([]byte, inputLength)
	for i := range input {
		input[i] = byte(rand.Intn(math.MaxUint8))
		if input[i] == wantedByte {
			occurencies++
		}
	}
}

func BenchmarkCountByChannels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if countByteOccurencesChannel(input, wantedByte) != occurencies {
			b.Fatalf("invalid occurencies count %d", occurencies)
		}
	}
}

func BenchmarkCountBySyncConcurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if countByteOccurencesSyncConcurrency(input, wantedByte) != occurencies {
			b.Fatalf("invalid occurencies count %d", occurencies)
		}
	}
}

func BenchmarkCountBySync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if countByteOccurencesSync(input, wantedByte) != occurencies {
			b.Fatalf("invalid occurencies count %d", occurencies)
		}
	}
}
