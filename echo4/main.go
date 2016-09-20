package main

import (
	"flag"
	"fmt"
)

var n = flag.Uint("n", 1, "count")
var sep = flag.String("sep", "\n", "separator")
var str = flag.String("str", "str", "string")

func main() {
	flag.Parse()
	var result string
	var i uint
	for ; i < *n; i++ {
		result = result + *str + *sep
	}

	fmt.Println(result)
}
