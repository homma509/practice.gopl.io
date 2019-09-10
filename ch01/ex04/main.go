package main

import (
	"bufio"
	"fmt"
	"os"
)

type count struct {
	n        int
	filename string
}

func main() {
	counts := make(map[string]*count)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup02: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, count := range counts {
		if count.n > 1 {
			fmt.Printf("%d\t%s\t%s\n", count.n, line, count.filename)
		}
	}
}

func countLines(f *os.File, counts map[string]*count) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if val, ok := counts[input.Text()]; ok {
			val.n++
		} else {
			counts[input.Text()] = &count{n: 1, filename: f.Name()}
		}
	}
}
