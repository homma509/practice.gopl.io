package main

import (
	"fmt"
	"io"
	"os"
)

// Echo args with space
func Echo(out io.Writer, args []string) {
	for i, arg := range args {
		fmt.Fprintln(out, i, arg)
	}
}

func main() {
	Echo(os.Stdout, os.Args[1:])
}
