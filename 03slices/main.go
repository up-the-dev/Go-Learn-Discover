package main

import "fmt"

func main() {
	fmt.Println("Slices in go")

	var list = []int{}
	list = append(list, 10, 20, 40, 04)
	fmt.Println(list)
	list = append(list[1:])
	fmt.Println(list)
	list = append(list[1:len(list)])

	list2 := []int{12, 3, 4, 9, 2, 3}
	list = list2
	fmt.Println(list)

	//removing element from slice based on index
	removableIndex := 1
	list = append(list[:removableIndex], list[removableIndex+1:]...)
	fmt.Println("list", list)
	fmt.Println("list2", list2)

	//POC of how capacity in slice behaves differently on how you create it.

	/* scenario 1: create slice using below syntax of using make */
	list3 := []int{2, 31, 45}
	fmt.Println(list3, len(list3), cap(list3))

	list3 = append(list3, 2)
	fmt.Println(list3, len(list3), cap(list3))

	list3 = append(list3, 2)
	fmt.Println(list3, len(list3), cap(list3))

	list3 = append(list3, 2, 4)
	fmt.Println(list3, len(list3), cap(list3))

	/* scenario 1: create slice using previous declared array */
	arr := [...]int{13, 23, 1}
	fmt.Println(arr, len(arr), cap(arr))
	slice := arr[0:len(arr)]
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))

	slice3 := make([]int, 3, 4)
	fmt.Println("checking the value of element which is out of length:", slice3)

	slice4 := [][][]string{
		{{"a1"}, {"a2", "a3"}, {"a4", "a5", "a6"}},
		{{"b1"}},
	}
	fmt.Println("multidimentional slice:", slice4)

	slc := []int{12, 45, 21, 43}

	fmt.Println("before modify first: ", slc, len(slc), cap(slc))
	modifyFirst(slc)

	fmt.Println("after modify first: ", slc, len(slc), cap(slc))

}

func modifyFirst(slice []int) {
	slice[0] = 1000

	slice = append(slice, 34, 343, 232, 2)
}
