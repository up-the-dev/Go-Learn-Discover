package main

import "fmt"

func addFROMFunc(add func(...int) int) func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {
	funcAsParam := func(slc ...int) int {
		sum := 0
		for _, v := range slc {
			sum = sum + v
		}
		return sum
	}
	outputFunc := addFROMFunc(funcAsParam)
	fmt.Println(outputFunc(2))
}
