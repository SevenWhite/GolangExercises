// exercise 1.8
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Request error:", err)
			os.Exit(1)
		}

		written, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintln(os.Stderr, "ReadAll error:", err)
			os.Exit(1)
		}

		fmt.Println("Written", written)
	}
}
