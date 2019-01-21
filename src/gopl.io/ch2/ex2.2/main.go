package main

import (
	"fmt"
	"os"
	"strconv"

	tempconv "gopl.io/ch2/ex2.1"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s = %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
		fmt.Printf("%s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
		fmt.Printf("%s = %s = %s\n", k, tempconv.KToF(k), tempconv.KToC(k))
	}
}
