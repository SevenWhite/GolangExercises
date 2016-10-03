package main

import (
	"bufio"
	"fmt"
)

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	advance, token, err := bufio.ScanWords(p, true)

	if err != nil {
		return advance, err
	}

	fmt.Println(advance, token, err)

	*c += WordsCounter(advance)
	return advance, nil
}

func main() {
	sentence := []byte("one two three four")
	var wc WordsCounter
	wc.Write(sentence)
	fmt.Println("wc", wc)
}
