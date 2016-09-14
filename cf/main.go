package main

import (
	"fmt"
	"github.com/SevenWhite/GolangExercises/tempconv"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args[1:]

	for _, number := range numbers {
		number, err := strconv.ParseFloat(number, 64)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Incorrect number", err)
			continue
		}

		c := tempconv.Celsius(number)
		f := tempconv.Fahrenheit(number)
		k := tempconv.Kelvin(number)

		fmt.Printf("%s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
		fmt.Printf("%s = %s = %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
		fmt.Printf("%s = %s = %s\n", k, tempconv.KToC(k), tempconv.KToF(k))
	}
}
