package main

import (
	"fmt"
	. "structs/backend"
	. "structs/user"
)

func main() {
	misha := NewUser("Миша", "misha@example.com")
	zheka := NewUser("Женя", "zheka@example.com")
	users := []*User{misha, zheka}
	for _, user := range users {
		PrintUserInfo(user)
	}
	ordersForToday := GenerateRandomOrderSlice(users, 10)
	paidOrdersForToday := OnlyPaidOrders(ordersForToday)
	PrintOrdersInfo(paidOrdersForToday)
	cashbackForToday := CashBackCalculatorWithCondition(5, 50) //кэшбек на сегодняшний набор покупок

	for _, order := range paidOrdersForToday {
		cash := cashbackForToday(order.Amount)
		order.User.Cashback += cash
	}
	PrintUserInfo(misha)
	misha.ApplyCashback()
	PrintUserInfo(misha)
	fmt.Println("В общем Жека оплатил заказов на сумму:", TotalUsersOrdersAmount(zheka, ordersForToday))
	TopSpenders(ordersForToday, 5)

}
