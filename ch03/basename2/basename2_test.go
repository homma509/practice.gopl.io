package main

import "testing"

func TestBasename(t *testing.T) {
	testCases := []struct {
		name string
		path string
		base string
		f    func(string) string
	}{
		{"basename2", "a", "a", basename},
		{"basename2", "a.go", "a", basename},
		{"basename2", "a/b/c.go", "c", basename},
		{"basename2", "a/b.c.go", "b.c", basename},
	}

	for _, c := range testCases {
		if c.f(c.path) != c.base {
			t.Errorf("%s: %s of base is wrong(%s)\n", c.name, c.path, c.f(c.path))
		}
	}
}
