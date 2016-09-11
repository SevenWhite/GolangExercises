package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Request error:", err)
			os.Exit(1)
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintln(os.Stderr, "ReadAll error:", err)
			os.Exit(1)
		}

		fmt.Printf("%s", bytes)
	}
}
