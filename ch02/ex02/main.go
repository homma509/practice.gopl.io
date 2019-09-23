package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/hommakoki/practice.gopl.io/ch02/ex02/conv"
)

const (
	defaultUnit = "templature"
)

var (
	u string
)

func main() {
	flag.StringVar(&u, "unit", defaultUnit, "convert unit type")
	flag.StringVar(&u, "u", defaultUnit, "convert unit type(short)")
	flag.Parse()

	args := make([]string, flag.NArg())
	copy(args, flag.Args())
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			args = append(args, input.Text())
		}
	}
	for _, arg := range args {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
			os.Exit(1)
		}
		unitconv(u, v)
	}
}

func unitconv(u string, v float64) {
	if u == "templature" {
		convtemp(v)
	}
	if u == "length" {
		convlength(v)
	}
	if u == "weight" {
		convweight(v)
	}
}

func convtemp(v float64) {
	f := conv.Fahrenheit(v)
	c := conv.Celsius(v)
	fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
}

func convlength(v float64) {
	m := conv.Meter(v)
	f := conv.Feet(v)
	fmt.Printf("%s = %s, %s = %s\n", m, conv.MToFt(m), f, conv.FtToM(f))
}

func convweight(v float64) {
	k := conv.Kilogram(v)
	p := conv.Pound(v)
	fmt.Printf("%s = %s, %s = %s\n", k, conv.KToP(k), p, conv.PToK(p))
}
