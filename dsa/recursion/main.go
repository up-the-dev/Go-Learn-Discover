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

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// fmt.Println(FindGreatestUsingLoop([]int{1, 4, 6, 5}))
	// PrintName5Times()
	// printFrom1toN(5)
	// printFrom1toNBackTracking(5)
	// printFromNto1(5)
	// printFromNto1Backtracking(1, 5)
	// fmt.Println(factorial(5))

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func checkIfPallindrome(str string, ptr int) bool {
	if ptr >= len(str)/2 {
		return true
	}
	if str[ptr] != str[len(str)-1-ptr] {
		return false
	}
	return checkIfPallindrome(str, ptr+1)
}

func isPalindrome(s string) bool {
	// remove the non-alphanumeric characters
	filteredStr := ""
	for _, chr := range s {
		if (chr >= 48 && chr <= 57) || (chr >= 65 && chr <= 90) || (chr >= 97 && chr <= 122) {
			if chr >= 65 && chr <= 90 {
				filteredStr = filteredStr + string(chr+32)
			} else {
				filteredStr = filteredStr + string(chr)

			}
		}
	}

	//check if isPalindrome
	return checkIfPallindrome(filteredStr, 0)

}
