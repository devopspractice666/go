package main

import "fmt"

func FilterNumbers(args ...int) (even, odd []int) {
	even = []int{}
	odd = []int{}
	for _, number := range args {
		if number%2 == 0 {
			even = append(even, number)
		} else {
			odd = append(odd, number)
		}
	}
	return
}

func main() {
	even, odd := FilterNumbers(1, 55, 1, 2, 6, 88, 9)
	fmt.Println("Четные:", even)
	fmt.Println("Нечетные", odd)
	numberSlice := []int{2, 666, 23, 32, 1, 8, 0, 54, 53}
	numberSlice2 := []int{3, 23, 17, 18, 143}
	numberSlice = append(numberSlice, numberSlice2...)
	even, odd = FilterNumbers(numberSlice...)
	fmt.Println("Четные:", even)
	fmt.Println("Нечетные", odd)
}

//упаковка в слайс в функции с несколькими параметрами
//распаковка слайса когда передается как аргумент в функцию
//полезная фича с объединением слайсов через распаковку
