package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Start")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Launch")
	case <-abort:
		fmt.Println("Launch aborted!")
	}
}
