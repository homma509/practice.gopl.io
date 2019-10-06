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
		{"comma", "1.123456", "1.123456", comma},
		{"comma", "12.123456", "12.123456", comma},
		{"comma", "123.123456", "123.123456", comma},
		{"comma", "1234.123456", "1,234.123456", comma},
		{"comma", "12345.123456", "12,345.123456", comma},
		{"comma", "123456.123456", "123,456.123456", comma},
		{"comma", "1234567.123456", "1,234,567.123456", comma},
		{"comma", "12345678.123456", "12,345,678.123456", comma},
		{"comma", "+1", "+1", comma},
		{"comma", "+12", "+12", comma},
		{"comma", "+123", "+123", comma},
		{"comma", "+1234", "+1,234", comma},
		{"comma", "+12345", "+12,345", comma},
		{"comma", "+123456", "+123,456", comma},
		{"comma", "+1234567", "+1,234,567", comma},
		{"comma", "+12345678", "+12,345,678", comma},
		{"comma", "+1.123456", "+1.123456", comma},
		{"comma", "+12.123456", "+12.123456", comma},
		{"comma", "+123.123456", "+123.123456", comma},
		{"comma", "+1234.123456", "+1,234.123456", comma},
		{"comma", "+12345.123456", "+12,345.123456", comma},
		{"comma", "+123456.123456", "+123,456.123456", comma},
		{"comma", "+1234567.123456", "+1,234,567.123456", comma},
		{"comma", "+12345678.123456", "+12,345,678.123456", comma},
		{"comma", "-1", "-1", comma},
		{"comma", "-12", "-12", comma},
		{"comma", "-123", "-123", comma},
		{"comma", "-1234", "-1,234", comma},
		{"comma", "-12345", "-12,345", comma},
		{"comma", "-123456", "-123,456", comma},
		{"comma", "-1234567", "-1,234,567", comma},
		{"comma", "-12345678", "-12,345,678", comma},
		{"comma", "-1.123456", "-1.123456", comma},
		{"comma", "-12.123456", "-12.123456", comma},
		{"comma", "-123.123456", "-123.123456", comma},
		{"comma", "-1234.123456", "-1,234.123456", comma},
		{"comma", "-12345.123456", "-12,345.123456", comma},
		{"comma", "-123456.123456", "-123,456.123456", comma},
		{"comma", "-1234567.123456", "-1,234,567.123456", comma},
		{"comma", "-12345678.123456", "-12,345,678.123456", comma},
	}
	for _, c := range testCases {
		if c.f(c.value) != c.format {
			t.Errorf("%s: %s of comma is wrong(expected=%s, actual=%s)", c.name, c.value, c.format, c.f(c.value))
		}
	}
}
