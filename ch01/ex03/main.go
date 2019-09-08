package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func EchoCat(out io.Writer, args ...string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)

}

func EchoBuf(out io.Writer, args ...string) {
	var b bytes.Buffer
	for i, s := range args {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(s)
	}
	fmt.Fprintln(out, b.String())
}

func EchoJoin(out io.Writer, args ...string) {
	fmt.Fprintln(out, strings.Join(args, " "))
}

func main() {
	EchoCat(os.Stdout, os.Args[1:]...)
}
