package main

import "fmt"

func PrintNameNtimes(counter int) {
	if counter == 0 {
		return
	}
	PrintNameNtimes(counter - 1)
	fmt.Println(counter)
}

func main() {
	PrintNameNtimes(5)
}
