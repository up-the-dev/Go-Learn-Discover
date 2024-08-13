package main

import "fmt"

func append1(list []int) {
	list = append(list, 4)
	list[0] = 34
	fmt.Println("list inside append1", list)
}

func main() {
	list := []int{1, 45, 2}
	fmt.Println("list inside main before func pass", list)
	append1(list)
	fmt.Println("list inside main after func pass", list)
}
