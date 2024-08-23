// this is ideal implementation of what we want to achieve in 05structs/paymentImpl.go where we did this using Structs
package main

import "fmt"

type Paymentor interface {
	Pay(amount float64)
}

type Payment struct {
	Paymentor
}

func (P Payment) makePayment(amount float64) {
	P.Paymentor.Pay(amount)
}

type Stripe struct{}

func (S Stripe) Pay(amount float64) {
	fmt.Println("amount of rs", amount, "is paid")
}
func main() {
	stripePG := Stripe{}
	newPayment := Payment{
		Paymentor: stripePG,
	}
	newPayment.makePayment(800)

}
