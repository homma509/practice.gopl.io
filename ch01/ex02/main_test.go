package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	out := new(bytes.Buffer)
	Echo(out, []string{"go", "programing", "language"})

	got := out.String()
	want := "0 go\n1 programing\n2 language\n"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
