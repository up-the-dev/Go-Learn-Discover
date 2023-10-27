package main

import "fmt"

/*
Short hand declaration not allowed outside function
age := 10
fmt.Printf("age: %d", age)
*/

const I_am_going_to_export_so_my_name_started_with_Uppercase_character string = "testVal"

func main() {
	var userID int = 10
	fmt.Println(userID)

	var isMale = true
	fmt.Println(isMale)

	var floatNum float32 = 122.32434234234234234
	fmt.Println(floatNum)

	var userName string = "umesh pawar"
	fmt.Printf("type is %T \n", userName)

	age := 10
	fmt.Printf("age: %d \n", age)

	//default value check

	var num int
	fmt.Println("int default value is :", num)

	var isNum bool
	fmt.Println(isNum)

	var numFloat float32
	fmt.Println(numFloat)

	var numString string
	fmt.Println("string default value is :", numString == "")

}
