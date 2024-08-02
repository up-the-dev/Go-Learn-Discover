package main

import "fmt"

//you generally use to export these struct and use its properties that why it's name and its key names are declared by starting capital letter
type User struct {
	Name   string
	Age    int
	Status bool
}

func main() {
	fmt.Println("structs in golang")
	mainUser := User{"umesh", 22, true}
	fmt.Println(mainUser)
	fmt.Printf("details user: %+v\n", mainUser)
	fmt.Println("name:", mainUser.Name)
}
