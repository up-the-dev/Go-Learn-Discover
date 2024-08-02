package main

import "fmt"

func main() {
	loginCount := 34

	if loginCount > 3 {
		fmt.Println("your", loginCount, "logins were done.")
		return
	}
	fmt.Println("logged in succesfully!")
}
