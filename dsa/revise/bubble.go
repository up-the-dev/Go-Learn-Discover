/*
adjacent swapping
{ 12,3,45,2,3,1,32} --> { 3,12,45,2,3,1,32} --> { 3,12,45,2,3,1,32} --> {3,12,2,45,3,1,32} --> {3,12,2,3,45,1,32}--> {3,12,2,3,1,32,45}

*/

package main

import "fmt"

func main() {
	arr := []int{12, 3, 45, 2, 65, 1, 32}

	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
