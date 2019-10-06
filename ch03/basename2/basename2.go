package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s = %s\n", arg, basename(arg))
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
