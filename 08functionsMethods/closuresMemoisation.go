package main

import "fmt"

func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(x int) int {
		if val, found := cache[x]; found {
			return val
		}
		result := f(x)
		cache[x] = result
		return result
	}
}

func slowFunction(x int) int {
	// Simulate a slow computation
	return x * x
}

func main() {
	memoizedFunc := memoize(slowFunction)

	fmt.Println(memoizedFunc(5)) // Computed, Output: 25
	fmt.Println(memoizedFunc(5)) // Cached, Output: 25
}
