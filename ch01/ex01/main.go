package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Echo args with space
func Echo(out io.Writer, args []string) {
	fmt.Fprintln(out, strings.Join(args, " "))
}

func main() {
	Echo(os.Stdout, os.Args[0:])
}
