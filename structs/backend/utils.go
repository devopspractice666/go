package backend

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	. "structs/order"
	. "structs/user"
)

func OnlyPaidOrders(ord []*Order) []*Order {
	var paidOrders []*Order
	for _, order := range ord {
		if order.Status == "paid" {
			paidOrders = append(paidOrders, order)
		}
	}
	return paidOrders
}
func FilterOrdersByStatus(orders []*Order, status string) []*Order {
	var needOrders []*Order
	for _, order := range orders {
		if order.Status == status {
			needOrders = append(needOrders, order)
		}
	}
	return needOrders

}
func OrderStats(orders []*Order) map[string]int {
	ord := make(map[string]int)
	for _, order := range orders {
		ord[order.Status]++
	}
	return ord
}

func GenerateRandomOrderSlice(userSLice []*User, size int) []*Order {
	orders := make([]*Order, 0, size)
	statuses := []string{"pending", "declined", "paid"}
	for i := 0; i < size; i++ {
		order := NewOrder(userSLice[rand.Intn(len(userSLice))], float64(rand.Intn(10666))/100)
		order.Status = statuses[rand.Intn(3)]
		orders = append(orders, order)
	}
	return orders
}
func PrintOrdersInfo(orders []*Order) {
	for _, order := range orders {
		fmt.Println("\nЗаказ номер:", order.ID)
		fmt.Print(strings.Repeat("-", 50))
		fmt.Println("\nУникальный номер заказчика :", order.User.ID, "\nИмя заказчика:", order.User.Name, "\nСтатус заказа:", order.Status)
		fmt.Println("Стоимость заказа", order.Amount)
		fmt.Print("\n")
	}
}

func TotalUsersOrdersAmount(u *User, orders []*Order) float64 {
	var total float64
	for _, order := range orders {
		if u.ID == order.User.ID && order.Status == "paid" {
			total += order.Amount
		}
	}
	return total

}
func TopSpenders(orders []*Order, n int) {
	type spending struct {
		user   *User
		amount float64
	}
	o := OnlyPaidOrders(orders)
	us := make(map[*User]float64)
	for _, order := range o {
		us[order.User] += order.Amount
	}
	var slice []*spending
	for key, value := range us {
		slice = append(slice, &spending{user: key, amount: value})
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].amount > slice[j].amount
	})

	fmt.Println("Топ пользователей по покупкам:")
	limit := n
	if len(slice) < limit {
		limit = len(slice)
	}
	for i := 0; i < limit; i++ {
		fmt.Println("Пользователь:", slice[i].user.Name, " Сумма:", slice[i].amount)
	}
}
