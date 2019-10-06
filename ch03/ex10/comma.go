package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s=%s\n", arg, comma(arg))
	}
}

func comma(s string) string {
	var buf bytes.Buffer

	for i, v := range []byte(s) {
		m := (len(s) - i) % 3
		if m == 0 && buf.Len() > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(v)
	}
	return buf.String()
}
