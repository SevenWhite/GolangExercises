package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%s", r.URL.Path)
}
