package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Balance  float64
	Cashback float64
}
type Order struct {
	ID     int
	User   *User
	Amount float64
	Status string
}

func idGenerator() func() int {
	id := 1
	return func() int {
		current := id
		id++
		return current
	}
}

var newID = idGenerator()
var NextOrderID = idGenerator()

func NewUser(name, email string) *User {
	return &User{
		ID:       newID(),
		Name:     name,
		Email:    email,
		Balance:  0,
		Cashback: 0,
	}
}
func NewOrder(user *User, amount float64) *Order {
	return &Order{
		ID:     NextOrderID(),
		User:   user,
		Amount: amount,
		Status: "pending",
	}
}
func (u *User) applyCashback() float64 {
	var amount float64 = u.Cashback
	u.Balance, u.Cashback = u.Balance+u.Cashback, 0
	return amount
}
func cashBackCalculator(percent float64) func(float64) float64 { //замыкание фукнции = анонимка + окружение(здесь это процент)
	return func(amount float64) float64 {
		return percent * amount / 100
	}
}
func cashBackCalculatorWithCondition(percent float64, minAmount float64) func(float64) float64 { //замыкание фукнции = анонимка + окружение(здесь это процент)
	return func(amount float64) float64 {
		if amount > minAmount {
			return percent * amount / 100
		}
		return 0
	}
} //улучшенная версия кэшбека по минимальной сумме заказа
func onlyPaidOrders(ord []*Order) []*Order {
	var paidOrders []*Order
	for _, order := range ord {
		if order.Status == "paid" {
			paidOrders = append(paidOrders, order)
		}
	}
	return paidOrders
}
func generateRandomOrderSlice(userSLice []*User, size int) []*Order {
	orders := make([]*Order, 0, size)
	statuses := []string{"pending", "declined", "paid"}
	for i := 0; i < size; i++ {
		order := NewOrder(userSLice[rand.Intn(len(userSLice))], float64(rand.Intn(10666))/100)
		order.Status = statuses[rand.Intn(3)]
		orders = append(orders, order)
	}
	return orders
}
func printOrdersInfo(orders []*Order) {
	for _, order := range orders {
		fmt.Println("\nЗаказ номер:", order.ID)
		fmt.Print(strings.Repeat("-", 50))
		fmt.Println("\nУникальный номер заказчика :", order.User.ID, "\nИмя заказчика:", order.User.Name, "\nСтатус заказа:", order.Status)
		fmt.Println("Стоимость заказа", order.Amount)
		fmt.Print("\n")
	}
}

func main() {
	misha := NewUser("Миша", "misha@example.com")
	zheka := NewUser("Женя", "zheka@example.com")
	users := []*User{misha, zheka}
	ordersForToday := generateRandomOrderSlice(users, 10)
	paidOrdersForToday := onlyPaidOrders(ordersForToday)
	printOrdersInfo(paidOrdersForToday)
	cashbackForToday := cashBackCalculatorWithCondition(5, 50) //кэшбек на сегодняшний набор покупок
	fmt.Println("Кэшбек Мишани до подсчета кэшбека:", misha.Cashback)
	for _, order := range paidOrdersForToday {
		cash := cashbackForToday(order.Amount)
		order.User.Cashback += cash
	}
	fmt.Println("Кэшбек Мишани после подсчета кэшбека:", misha.Cashback)
	// анонимная функция
	// paidOrders := func(orders []*Order) []*Order {

	// 	var paidOrd []*Order
	// 	for _, ord := range orders {
	// 		if ord.Status == "paid" {
	// 			paidOrd = append(paidOrd, ord)
	// 		}
	// 	}
	// 	return paidOrd
	// }(orders)

}
