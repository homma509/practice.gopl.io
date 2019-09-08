package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	out := new(bytes.Buffer)
	Echo(out, []string{"./echo", "go", "programing", "language"})

	got := out.String()
	want := "./echo go programing language\n"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
