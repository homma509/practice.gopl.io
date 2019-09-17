package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func seed() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}

func bench(b *testing.B, f func(uint64) int) {
	// b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed())
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchCases := []struct {
		name string
		f    func(uint64) int
	}{
		{"OneLiner", PopCountOneLiner},
		{"BitLoop", PopCountBitLoop},
	}
	for _, c := range benchCases {
		b.Run(fmt.Sprintf("%s", c.name),
			func(b *testing.B) { bench(b, c.f) })
	}
}
