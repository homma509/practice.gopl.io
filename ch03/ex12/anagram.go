package main

import (
	"sort"
	"strings"
)

func anagram(s1, s2 string) bool {
	v1 := strings.Split(s1, "")
	v2 := strings.Split(s2, "")
	sort.Strings(v1)
	sort.Strings(v2)
	w1 := strings.Join(v1, "")
	w2 := strings.Join(v2, "")
	return strings.Contains(w1, w2) || strings.Contains(w2, w1)
}
