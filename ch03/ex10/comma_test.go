package main

import "testing"

func TestComma(t *testing.T) {
	testCases := []struct {
		name   string
		value  string
		format string
		f      func(string) string
	}{
		{"comma", "1", "1", comma},
		{"comma", "12", "12", comma},
		{"comma", "123", "123", comma},
		{"comma", "1234", "1,234", comma},
		{"comma", "12345", "12,345", comma},
		{"comma", "123456", "123,456", comma},
		{"comma", "1234567", "1,234,567", comma},
		{"comma", "12345678", "12,345,678", comma},
	}
	for _, c := range testCases {
		if c.f(c.value) != c.format {
			t.Errorf("%s: %s of comma is wrong(expected=%s, actual=%s)", c.name, c.value, c.format, c.f(c.value))
		}
	}
}
