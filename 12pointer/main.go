package main

import "fmt"

func mul10(num int) {
	num = num * 10
}
func mul10PTR(num *int) {
	*num = *num * 10
}

func main() {
	// null pointer
	var nilptr *int
	fmt.Println("value of nilPtr-->", nilptr)

	num := 18
	// pointer declaration
	var numptr *int = &num
	fmt.Println("address of num", numptr)
	fmt.Println("address of numptr", &numptr)
	fmt.Println("value of num", *numptr)

	// creating ptr using new
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)

	// handle variables value by Dereferencing its pointer
	*size = *size + 20
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)

	// passing by value vs passing by address in function
	fmt.Printf("num value is  before passing anywhere %d\n", num)
	mul10(num)
	fmt.Printf("num value after passing by value is %d\n", num)
	mul10PTR(numptr) /* or mul10PTR(&num) */
	fmt.Printf("num value after passing by address(ptr) is %d\n", num)

	// to make changes in array (you want to pass by address) then dont use pointre their use slice because slice is also referencing to underlying array by its address

}

// is nilptr in above code will garbage collected ?
/* In your code, nilptr is just a pointer variable that has been declared but not initialized to point to any valid memory address. Since it doesn't point to any allocated memory, there's no memory to reclaim, and it will not be garbage collected. Garbage collection in Go typically deals with dynamically allocated memory that is no longer in use.

To summarize, nilptr is not eligible for garbage collection because it doesn't point to any memory that needs to be managed by the garbage collector. It's just a pointer variable initialized to nil. */
