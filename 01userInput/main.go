package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("please enter rate go out of 10 ")

	reader := bufio.NewReader(os.Stdin)

	//this is comma,ok or err,ok syntax which is generally used like try catch in go.
	rating, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating : ", rating)
	fmt.Printf("type of input : %T", rating)

}
