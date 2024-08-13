/* Selection Sort algorithm
finds the lowest value in an array and moves it to the front of the array.

steps:
1.Go through the array to find the lowest value.
2.Move the lowest value to the front of the unsorted part of the array.
3.Go through the array again as many times as there are values in the array
*/

package main

// import "fmt"

// func main() {
// 	unsortedSLC := []int{64, 34, 25, 5, 22, 11, 90, 12}

// 	for i := 0; i < len(unsortedSLC)-1; i++ {
// 		min := i
// 		for j := i; j < len(unsortedSLC); j++ {
// 			if unsortedSLC[j] < unsortedSLC[min] {
// 				min = j
// 			}
// 		}
// 		if i != min {
// 			// swaping unsortedSLC[i] with unsortedSLC[min]
// 			temp := unsortedSLC[i]
// 			unsortedSLC[i] = unsortedSLC[min]
// 			unsortedSLC[min] = temp
// 		}

// 		fmt.Println(unsortedSLC)

// 	}
// 	fmt.Println("sortedSLC:", unsortedSLC)

// }
