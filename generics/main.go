package main

import "fmt"

func interator[T int | string](slc []T) {
	fmt.Printf("type: %T \n", slc)
	for _, v := range slc {
		fmt.Println(v)
	}
}

// here we duplicated same code just because we wat to use this behaviour for different type which is wrong practice(DRY - do not repeat yourself)
func interatorstrings(slc []string) {
	fmt.Printf("type: %T \n", slc)
	for _, v := range slc {
		fmt.Println(v)
	}
}

type Stack[T comparable] struct {
	elements []T
}

func main() {
	// interator([]int{12, 34, 5})
	// interator([]string{"umes", "paew"})

	myStack := Stack[int]{
		[]int{12, 34, 45},
	}

	myStack2 := Stack[string]{
		[]string{"umesh", "pawar"},
	}
	fmt.Println(myStack, myStack2)
}
