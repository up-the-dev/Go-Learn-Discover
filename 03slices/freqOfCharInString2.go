/*
input :"rutuja"
output: { r:1,u:2,t:1,j:1,a:1}
*/
package main

import "fmt"

func convertStrToMapOfRunes(inputstr string) map[rune]int {
	mapOfRunes := make(map[rune]int)
	for _, v := range inputstr {
		mapOfRunes[v] = 0
	}

	return mapOfRunes
}

func main() {
	inputStr := "rutuja"
	inputMap := convertStrToMapOfRunes(inputStr)

	for _, v := range inputStr {
		_, ok := inputMap[v]
		if ok {
			inputMap[v]++
		}
	}

	fmt.Println(inputMap)
}
