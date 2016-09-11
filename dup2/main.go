package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) > 0 {
		readFiles(files, counts)
	} else {
		countLines(os.Stdin, counts)
	}

	for str, count := range counts {
		if count > 1 {
			fmt.Printf("%s(%d)\n", str, count)
		}
	}
}

func readFiles(files []string, counts map[string]int) {
	for _, file := range files {
		file, err := os.Open(file)

		if err != nil {
			fmt.Fprintln(os.Stderr, "File open error: ", err)
			continue
		}

		countLines(file, counts)
		file.Close()
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
}
