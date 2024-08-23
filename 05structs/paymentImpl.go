package main

import "fmt"

type Payment struct {
	PaymentGateway Stripe
}

func (P Payment) makePayment(amount float64) {
	P.PaymentGateway.Pay(amount)
}

type Stripe struct{}

func (S Stripe) Pay(amount float64) {
	fmt.Println("payment of", amount, "is successfull.")
}

func main() {
	StripePG := Stripe{}
	newPayment := Payment{StripePG}
	newPayment.makePayment(1200)
}
