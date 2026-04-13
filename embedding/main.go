package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) speak() string {
	return "Animal speaks"
}

func (a *Animal) Info() string {
	info := "Animal " + a.Name
	return info
}

type Dog struct {
	Animal
	Breed string
}

func (d *Dog) speak() string {
	return "Woof"
}

func (d *Dog) Info() string {
	info := "Animal " + d.Name + " Breed:" + d.Breed
	return info
}

type Cat struct {
	Animal
	Color string
}

func (c *Cat) Info() string {
	info := "Animal " + c.Name + " Color:" + c.Color
	return info
}

type Speaker interface {
	speak() string
	Info() string
}

func MakeSound(s Speaker) {
	fmt.Println(s.speak())
}

func main() {
	animal := Animal{Name: "Just Animal"}
	dog := Dog{Animal: Animal{Name: "Sharik"}, Breed: "German Shepherd"}
	cat := Cat{Animal: Animal{Name: "Murzik"}, Color: "Red"}

	slice := []Speaker{&animal, &dog, &cat}

	for _, element := range slice {
		fmt.Println(element.Info())
		MakeSound(element)
	}

}

//работа с интерфейсами, встраивание структур, реализация методов интерфейса
//методы dog затеняют методы из родительской структуры
//кот юзает метод speak от animal т.к не реализован свой
//создан метод принимающий объект интерфейса, который выводит спик для каждой стркутуры свой
