package main

import (
	"fmt"
)

type EmployeeSal struct {
	name   string
	salary float64
}

func main() {
	fmt.Println("maps in golang")
	langs := make(map[string]string)
	langs["rb"] = "ruby"
	langs["go"] = "goLang"
	langs["js"] = "javascript"
	langs["py"] = "python"
	fmt.Println(langs)
	//to remove from map use delete
	delete(langs, "py")
	fmt.Println(langs)

	fmt.Println("example 2")
	emp1 := EmployeeSal{
		name:   "joe",
		salary: 1123.11112,
	}

	emp2 := EmployeeSal{
		name:   "ram",
		salary: 20000.45,
	}

	salDb := map[string]EmployeeSal{
		"joeID": emp1,
		"ramID": emp2,
	}

	for emplId, employee := range salDb {
		fmt.Printf("salDb[%s]= { name: %s, salary: %f}\n", emplId, employee.name, employee.salary)
	}

	map_using_make := make(map[string]int, 4)
	fmt.Println(map_using_make, len(map_using_make))

	marks1 := [5]int{11, 13, 14, 9, 18}
	marks2 := [5]int{9, 01, 20, 15, 10}

	markDetailMap := map[[5]int]map[string]int{
		marks1: map[string]int{
			"class":      10,
			"totalMarks": 65,
		},
		marks2: map[string]int{
			"class":      9,
			"totalMarks": 55,
		},
	}

	iterator := 1
	for marksList, desc := range markDetailMap {

		fmt.Println("iteration : ", iterator)
		fmt.Println("key : ", marksList)

		for _, value := range marksList {
			fmt.Println(value)
		}
		for key, value := range desc {
			fmt.Println("key of desc: ", key)
			fmt.Println("value of desc: ", value)

		}
		iterator += 1
	}
	duplicateSlc := []int{23, 45, 2, 23, 2, 45}
	fmt.Println("duplicateSlc:", duplicateSlc)
	RemoveDuplicates(duplicateSlc)
	fmt.Println("uniqueSlc:", duplicateSlc)

}

func RemoveDuplicates(slcWithDuplicate []int) []int {
	// convert this slcWithDuplicate to a map which should have keys with slcWithDuplicate's elements
	set := make(map[int]struct{}, len(slcWithDuplicate))

	for _, value := range slcWithDuplicate {
		set[value] = struct{}{}
	}

	// create a slice with same cap and len as slcWithDuplicates and push these unique key into it
	uniqueSlc := []int{}
	for key := range set {
		uniqueSlc = append(uniqueSlc, key)
	}
	return uniqueSlc
}
