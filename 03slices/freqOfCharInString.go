package main

import "fmt"

func main() {
	inputStr := "rutuja"
	inputRune := []rune(inputStr)

	// approch using slc
	freqSlc := make([]int32, 65536)

	for _, v := range inputRune {
		freqSlc[v] = freqSlc[v] + 1
	}
	for i, v := range freqSlc {
		if v != 0 {
			fmt.Println(i, v)
		}
	}

	//approch using map
	freqMap := make(map[rune]int)

	for _, v := range inputRune {
		freqMap[v] = freqMap[v] + 1
	}
	fmt.Println("freqmap: ", freqMap)
}
