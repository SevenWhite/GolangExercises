package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for word, count := range counts {
		if count > 1 {
			fmt.Printf("%s(%d) \n", word, count)
		}
	}
}
