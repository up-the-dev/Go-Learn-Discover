package main

import "fmt"

func main() {
	var nums = [5]int{10, 12, 34, 7, 86}
outer:
	for i, _ := range nums {
		for j := 0; j <= i; j++ {
			fmt.Println(i, j)
			if j == 1 {
				break outer
			}
		}
	}
}
