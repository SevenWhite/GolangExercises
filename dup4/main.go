// exercise 1.4
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Count struct {
	number    int
	filePaths map[string]int
}

type Counts map[string]Count

func main() {
	counts := make(Counts)
	files := os.Args[1:]

	if len(files) > 0 {
		readFiles(files, counts)
	} else {
		countLines(os.Stdin, counts, "os.Stdin")
	}

	for str, count := range counts {
		if count.number > 1 {
			fmt.Printf("%s(%d) files(%v)\n", str, count.number, count.filePaths)
		}
	}
}

func readFiles(files []string, counts Counts) {
	for _, filePath := range files {
		file, err := os.Open(filePath)

		if err != nil {
			fmt.Fprintln(os.Stderr, "File open error: ", err)
			continue
		}

		countLines(file, counts, filePath)
		file.Close()
	}
}

func countLines(file *os.File, counts Counts, filePath string) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		text := input.Text()
		count := counts[text]
		count.number++

		if len(count.filePaths) == 0 {
			count.filePaths = map[string]int{filePath: 1}
		} else {
			count.filePaths[filePath]++
		}

		counts[text] = count
	}
}
