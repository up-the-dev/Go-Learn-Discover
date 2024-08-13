package main

import (
	"fmt"
	"time"
)

func worker(id int) func() {
	count := 0
	return func() {
		for {
			count++
			fmt.Printf("Worker %d: %d\n", id, count)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	go worker(1)()
	go worker(2)()

	// Prevent main from exiting immediately
	select {}
}
