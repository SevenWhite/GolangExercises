package main

import (
	"flag"
	"fmt"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type CelsiusFlag struct {
	Celsius
}

func (f *CelsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	}
	return fmt.Errorf("Wrong temperature %q", s)
}

func CreateCelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := CelsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temperature = CreateCelsiusFlag("t", 20, "temperature (Celsius)")

func main() {
	flag.Parse()
	fmt.Printf("\nTemperature: %v", temperature)
}
