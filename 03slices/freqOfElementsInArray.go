// input arr:= []int{12,2,4,5,2,45,4,12} output : { 12:2, 2:2, 4:2,5:1,45:1 }
package main

import "fmt"

func main() {
	arr := []int{12, 2, 4, 5, 2, 45, 4, 12}
	freqMap := make(map[int]int)

	for _, v := range arr {
		_, ok := freqMap[v]
		if ok == true {
			freqMap[v]++
			continue
		}
		freqMap[v] = 1
	}

	fmt.Println("freqMap-->", freqMap)
}
