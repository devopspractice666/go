package backend

import (
	"math"
)

func CashBackCalculator(percent float64) func(float64) float64 { //замыкание фукнции = анонимка + окружение(здесь это процент)
	return func(amount float64) float64 {
		return percent * amount / 100
	}
}

func CashBackCalculatorWithCondition(percent float64, minAmount float64) func(float64) float64 { //замыкание фукнции = анонимка + окружение(здесь это процент)
	return func(amount float64) float64 {
		if amount > minAmount {
			result := percent * amount / 100
			return math.Round(result*100) / 100
		}
		return 0
	}
} //улучшенная версия кэшбека по минимальной сумме заказа
