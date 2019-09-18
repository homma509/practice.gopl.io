package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPopCount(t *testing.T) {
	testCases := []struct {
		name  string
		bits  uint64
		count int
		f     func(uint64) int
	}{
		{"PopCountOneLiner", 0, 0, PopCountOneLiner},
		{"PopCountOneLiner", 3, 2, PopCountOneLiner},
		{"PopCountOneLiner", 4, 1, PopCountOneLiner},
		{"PopCountOneLiner", 255, 8, PopCountOneLiner},
		{"PopCountBitClear", 0, 0, PopCountBitClear},
		{"PopCountBitClear", 3, 2, PopCountBitClear},
		{"PopCountBitClear", 4, 1, PopCountBitClear},
		{"PopCountBitClear", 255, 8, PopCountBitClear},
	}
	for _, c := range testCases {
		if got := c.f(c.bits); got != c.count {
			t.Errorf("%s: got: %v, want: %v", c.name, got, c.count)
		}
	}
}

func seed() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}

func bench(b *testing.B, f func(uint64) int) {
	b.ReportAllocs()
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
		{"BitClear", PopCountBitClear},
	}
	for _, c := range benchCases {
		b.Run(fmt.Sprintf("%s", c.name),
			func(b *testing.B) { bench(b, c.f) })
	}
}
