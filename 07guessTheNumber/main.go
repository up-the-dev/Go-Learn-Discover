package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("welcome to Guess the number game:")
	fmt.Println("rules:")
	fmt.Println("you have 5 chances to guess the number correctly.")
	randomGuessingWord := randGen.Intn(101)
	for i := 0; i < 5; i++ {
		fmt.Printf("Please enter a number : ")
		userInputStr, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("something went wrong while taking input :", err)
			return
		}
		userInputStr = strings.TrimSpace(userInputStr)
		userInput, err := strconv.Atoi(userInputStr)
		if err != nil {
			fmt.Println("something went wrong while converting input to string :", err)
			return
		}
		// fmt.Printf("%d %d", userInput, randomGuessingWord)
		if userInput == randomGuessingWord {
			fmt.Println("hurrey 🎇 ! you guessed it right 😊. number is :", randomGuessingWord)
			return
		} else if userInput > randomGuessingWord {
			fmt.Println("missed! you guessed bigger number 🙁. remaining trials:", 5-(i+1))
		} else if userInput < randomGuessingWord {
			fmt.Println("missed! you guessed smaller number 🙁. remaining trials:", 5-(i+1))
		}
	}
	fmt.Println("you lost 😩. number was :", randomGuessingWord)
}
