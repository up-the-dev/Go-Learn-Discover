package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Delivered
	Failed
)

func changeOrderStatus(status OrderStatus) {
	fmt.Printf("changing status to %d type is %t", status, status)
}

func main() {
	changeOrderStatus(Delivered)
}
