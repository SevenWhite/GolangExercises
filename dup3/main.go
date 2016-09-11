package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, filePath := range files {
		bytes, err := ioutil.ReadFile(filePath)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Read file error:", err)
			continue
		}

		stringsArr := strings.Split(string(bytes), "\r\n")

		for _, str := range stringsArr {
			counts[str]++
		}
	}

	for str, count := range counts {
		if count > 1 {
			fmt.Printf("%s(%d)\n", str, count)
		}
	}
}
