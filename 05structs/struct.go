package main

import (
	"fmt"
	"time"
)

// declaring a template
type Order struct {
	id          int
	name        string
	price       float64
	delieryTime time.Time
	time        time.Time
}

// constructor
func newOrder(id int, name string, price float64) *Order {
	return &Order{1, name, price, time.Now(), time.Now()}
}

// getter
func (o Order) GetDeliveryTime() time.Time {
	return o.delieryTime
}

// setter
func (o *Order) SetDeliveryTime(newDeliveryTime time.Time) {
	o.delieryTime = newDeliveryTime
}

func main() {

	// creating instance using constructor of struct
	order1 := newOrder(1, "watch", 1300.08)
	fmt.Println(order1)

	fmt.Println(order1.GetDeliveryTime())

	order1.SetDeliveryTime(time.Date(2024, 12, 31, 2, 45, 0, 0, time.UTC))
	fmt.Println(order1.GetDeliveryTime())
}
