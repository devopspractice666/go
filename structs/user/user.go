package user

import (
	"fmt"
	"strings"
	. "structs/generator"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Balance  float64
	Cashback float64
}

func NewUser(name, email string) *User {
	return &User{
		ID:       NewID(),
		Name:     name,
		Email:    email,
		Balance:  0,
		Cashback: 0,
	}
}

func (u *User) ApplyCashback() float64 {
	var amount float64 = u.Cashback
	u.Balance, u.Cashback = u.Balance+u.Cashback, 0
	return amount
}
func PrintUserInfo(u *User) {
	fmt.Println("Информация о пользователе с уникальным номером:", u.ID)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Имя:", u.Name, "\nПочта для связи:", u.Email, "\nБаланс пользователя:", u.Balance, " рублей")
	fmt.Println("КЭшбек пользователя:", u.Cashback, " рублей")
	fmt.Print("\n")
}
