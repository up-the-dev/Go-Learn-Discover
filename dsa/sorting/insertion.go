package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func main() {
	arr := []int{7, 12, 9, 11, 3}
	fmt.Println("unsorted array:", arr)
	insertionSort(arr)
	fmt.Println("sorted array:", arr)
}
