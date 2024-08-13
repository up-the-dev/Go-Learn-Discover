package main

import "fmt"

func filter(slice []int, test func(int) bool) []int {
    var result []int
    for _, v := range slice {
        if test(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6}

    isEven := func(n int) bool {
        return n%2 == 0
    }

    evens := filter(numbers, isEven)
    fmt.Println(evens) // Output: [2 4 6]
}
