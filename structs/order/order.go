package order

import (
	. "structs/generator"
	. "structs/user"
)

type Order struct {
	ID     int
	User   *User
	Amount float64
	Status string
}

func NewOrder(user *User, amount float64) *Order {
	return &Order{
		ID:     NextOrderID(),
		User:   user,
		Amount: amount,
		Status: "pending",
	}
}
