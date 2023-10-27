package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something crazy :")

	input, _ := reader.ReadString('\n')
	fmt.Println("your crazy input is: ", input)

}
