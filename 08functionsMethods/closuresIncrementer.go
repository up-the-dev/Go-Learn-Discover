package main

import "fmt"

func incrementer() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	inc1 := incrementer()
	inc2 := incrementer()

	fmt.Println(inc1()) // Output: 1
	fmt.Println(inc1()) // Output: 2
	fmt.Println(inc2()) // Output: 1
	fmt.Println(inc2()) // Output: 2
}
