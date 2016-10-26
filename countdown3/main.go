package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Start!")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("Abort!")
			return
		}
	}

	fmt.Println("Launch!")
}
