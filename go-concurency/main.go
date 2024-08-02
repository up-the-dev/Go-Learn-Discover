package main

import (
	"fmt"
)

func calculateDiscountInRS(amt, discnt *float64, chnl chan float64) {
	discountAmt := (*amt * *discnt) / 100

	chnl <- discountAmt
}

func main() {
	fmt.Printf("main function called !\n\n\n")
	totalAmt := new(float64)
	*totalAmt = 80
	discountInPercentage := new(float64)
	*discountInPercentage = 10
	chnl := make(chan float64)
	go calculateDiscountInRS(totalAmt, discountInPercentage, chnl)
	discountInRS := <-chnl
	fmt.Println("discountAmount calculated -->", discountInRS)

	fmt.Println("\n\nmain function terminated !")

}
