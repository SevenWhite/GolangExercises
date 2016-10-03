package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	var count int
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		count++
	}

	*c += LineCounter(count)
	return count, nil
}

func main() {
	lines := []byte("one\ntwo\nthree\nfour\nfive\nsix\n")

	var lc LineCounter
	lc.Write(lines)
	fmt.Println(string(lines))
	fmt.Println("lc", lc)
}
