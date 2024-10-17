/*
find min and move to first

{ 12,23,4,25,34}
{4,23,12,25,34}
{4,12,23,25,24}
{4,12,23,25,24}
{4,12,23,24,25}

1. find min swap with 1st unsorted element

*/

package main

import "fmt"

func main() {
	arr := []int{12, 23, 4, 25, 34, 2}

	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}
