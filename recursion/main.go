package main

import "fmt"

func FindGreatestUsingLoop(list []int) (greatest int) {
	greatest = list[0]
	for i := range list {
		if list[i] > greatest {
			greatest = list[i]
		}
	}
	return
}

func recursionPrint(num int) {
	fmt.Println(num)
	num++
	recursionPrint(num)
}

var count = 0

func PrintName5Times() {
	fmt.Println("umesh")
	count++
	if count == 5 {
		return
	}
	PrintName5Times()
}

func printFrom1toN(n int) {
	if n < counter {
		return
	}
	fmt.Println(counter)
	counter++
	printFrom1toN(n)

}

var counter = 1

func printFrom1toNBackTracking(n int) {
	if n == 0 {
		return
	}
	printFrom1toNBackTracking(n - 1)
	fmt.Println(n)

}

func printFromNto1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	n--
	printFromNto1(n)
}

func printFromNto1Backtracking(from int, to int) {
	if from > to {
		return
	}
	printFromNto1Backtracking(from+1, to)
	fmt.Println(from)
}

func main() {
	// fmt.Println(FindGreatestUsingLoop([]int{1, 4, 6, 5}))
	// PrintName5Times()
	// printFrom1toN(5)
	// printFrom1toNBackTracking(5)
	// printFromNto1(5)
	printFromNto1Backtracking(1, 5)
}
