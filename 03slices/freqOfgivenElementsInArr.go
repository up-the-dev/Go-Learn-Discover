/*
input

	searchInArr:= []int{12,2,4,5,2,45,4,12}
	elementsToBeSearch:= []int{2, 4, 5, 6, 12 ,0}

output : { 2:2, 4:2, 5:1, 6:0, 12:2, 0:0 }
*/
package main

import "fmt"

func ConvertToMap(elementsToBeSearchArr []int) map[int]int {
	elementToBeSearchMap := make(map[int]int)

	for _, v := range elementsToBeSearchArr {
		elementToBeSearchMap[v] = 0
	}

	return elementToBeSearchMap
}

func main() {
	searchInArr := []int{12, 2, 4, 5, 2, 45, 4, 12}
	elementsToBeSearch := []int{2, 4, 5, 6, 12, 0}
	elementToBeSearchMap := ConvertToMap(elementsToBeSearch)

	for _, v := range searchInArr {
		_, isVpresent := elementToBeSearchMap[v]
		if isVpresent {
			elementToBeSearchMap[v]++
		}
	}

	fmt.Printf("%d\n", elementToBeSearchMap)

}
