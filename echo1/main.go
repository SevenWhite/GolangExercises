// print command line arguments
package main

import (
	"fmt"
	"os"
)

const separator string = " "

func main() {
	var s string
	args := os.Args[1:]

	for _, arg := range args {
		s += separator + arg
	}

	fmt.Println(s)
}
