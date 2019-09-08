package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "test")
	}
	return s
}

func bench(b *testing.B, n int, f func(io.Writer, ...string)) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(ioutil.Discard, seed(n)...)
	}
}

func BenchmarkConcatenate(b *testing.B) {
	benchCases := []struct {
		name string
		n    int
		f    func(io.Writer, ...string)
	}{
		{"Cat", 3, EchoCat},
		{"Buf", 3, EchoBuf},
		{"Join", 3, EchoJoin},
		{"Cat", 100, EchoCat},
		{"Buf", 100, EchoBuf},
		{"Join", 100, EchoJoin},
		{"Cat", 10000, EchoCat},
		{"Buf", 10000, EchoBuf},
		{"Join", 10000, EchoJoin},
	}
	for _, c := range benchCases {
		b.Run(fmt.Sprintf("%s%d", c.name, c.n),
			func(b *testing.B) { bench(b, c.n, c.f) })
	}
}
