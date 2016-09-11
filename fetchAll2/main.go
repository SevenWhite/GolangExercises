// exercise 1.10
package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := os.Args[1:]

	for _, url := range urls {
		if !isValidUrl(url) {
			url = "http://" + url
		}

		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	fileName := fmt.Sprintf("%s%d.html", strings.Replace(url, "http://", "", -1), rand.Int())

	file, err := os.Create(fileName)

	if err != nil {
		ch <- fmt.Sprintf("Create file error %s: %v", url, err)
		return
	}

	nbyts, err := io.Copy(file, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbyts, url)
}

func isValidUrl(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
