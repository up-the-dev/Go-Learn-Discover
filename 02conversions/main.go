package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("please enter rating betwwn 1 to 5 :")
	rating, _ := reader.ReadString('\n')
	fmt.Printf("plane rating is : %s and type is %T \n", rating, rating)

	// so rating is of type string, now i want to perform ops on user enterd rating, for that i need to convert rating to number.

	//for string related conversion use strconv and for string ops use strings package
	ratingInNum, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("converted successfully ", ratingInNum)
	}
}
