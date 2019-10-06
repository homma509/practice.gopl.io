package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s=%s\n", arg, comma(arg))
	}
}

func comma(s string) string {
	n := len(s)

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		return s[:1] + comma(s[1:])
	}

	if i := strings.Index(s, "."); i >= 0 {
		return comma(s[:i]) + s[i:]
	}

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
