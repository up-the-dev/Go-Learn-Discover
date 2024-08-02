package main

import "fmt"

func main() {
	umesh := User{"umesh", 820870952, "umesh@gmail.com"}
	fmt.Println(umesh)
	fmt.Println("email:", umesh.email)
	/* 	funcSum := addNums
	   	fmt.Println(funcSum) */
	// fmt.Println(addNums(12, 12, 6, 6))
	updatedUserMail := umesh.SetUser("update@mail.in")
	fmt.Println(" updated mail:", updatedUserMail)
	fmt.Println("field updated :", umesh.email)
	fmt.Println(umesh)
}

type User struct {
	name   string
	mobile int
	email  string
}

func (u *User) SetUser(mail string) string {
	u.email = mail
	return u.email
}

func addNums(nums ...int) (string, int) {
	sum := 0

	for _, num := range nums {
		sum = sum + num
	}
	return "sum is", sum
}
