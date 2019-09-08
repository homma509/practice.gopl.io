package main

import (
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

func BenchmarkEchoCat3(b *testing.B)      { bench(b, 3, EchoCat) }
func BenchmarkEchoBuf3(b *testing.B)      { bench(b, 3, EchoBuf) }
func BenchmarkEchoJoin3(b *testing.B)     { bench(b, 3, EchoJoin) }
func BenchmarkEchoCat100(b *testing.B)    { bench(b, 100, EchoCat) }
func BenchmarkEchoBuf100(b *testing.B)    { bench(b, 100, EchoBuf) }
func BenchmarkEchoJoin100(b *testing.B)   { bench(b, 100, EchoJoin) }
func BenchmarkEchoCat10000(b *testing.B)  { bench(b, 10000, EchoCat) }
func BenchmarkEchoBuf10000(b *testing.B)  { bench(b, 10000, EchoBuf) }
func BenchmarkEchoJoin10000(b *testing.B) { bench(b, 10000, EchoJoin) }
