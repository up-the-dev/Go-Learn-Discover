package main

import (
	"fmt"
	"time"
)

func task(tasknumber int) {
	fmt.Println("task", tasknumber, "is running")
}

func main() {

	var wg 

	for i := 1; i <= 5; i++ {

		go task(i)
	}
	time.Sleep(2000000)
}
