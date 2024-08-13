package main

import "fmt"

func increment() func() int {
	var num int
	return func() int {
		num++
		return num
	}
}

func main() {
	increment1 := increment()
	fmt.Println(increment1())
	fmt.Println(increment1())
}
