package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	testCases := []struct {
		name    string
		s1      string
		s2      string
		anagram bool
		f       func(string, string) bool
	}{
		{"anagram", "a", "a", true, anagram},
		{"anagram", "abc", "cab", true, anagram},
		{"anagram", "a", "b", false, anagram},
		{"anagram", "abc", "bcdef", false, anagram},
	}

	for _, c := range testCases {
		if c.f(c.s1, c.s2) != c.anagram {
			t.Errorf("%s: %s, %s of anagram is wrong(expected=%v, actual=%v)", c.name, c.s1, c.s2, c.anagram, c.f(c.s1, c.s2))
		}
	}
}
