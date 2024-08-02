package main

import "fmt"

func main() {
	slc := []int{12, 45, 21, 43}
	fmt.Println("before modifyFirst: ", slc, len(slc), cap(slc))
	modifyFirst(slc)
	fmt.Println("after modifyFirst: ", slc, len(slc), cap(slc))
}

func modifyFirst(slice []int) {
	slice[0] = 1000
	slice = append(slice, 34, 343, 232, 2)
	slice[1] = 2000
}
