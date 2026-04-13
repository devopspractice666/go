package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() string {
	return "Animal speaks"
}

func (a *Animal) Info() string {
	info := "Animal " + a.Name
	return info
}

func (a *Animal) Move() string {
	return a.Name + " moves!"
}

type Dog struct {
	Animal
	Breed string
}

func (d *Dog) Speak() string {
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
	Speak() string
	Info() string
}

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

type Mover interface {
	Move() string
}

type AnimalInterface interface {
	Speaker
	Mover
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

	var test AnimalInterface = &Animal{Name: "Leo"}
	fmt.Println(test.Move())
	fmt.Println(test.Speak())
	fmt.Println(test.Info())

	var s Speaker = &Dog{Animal: Animal{Name: "Charlie"}, Breed: "Ovcharka"}
	puppy := s.(*Dog)
	println("Собака с именем ", puppy.Name)

	if dog2, ok := s.(*Cat); ok {
		fmt.Println(dog2.Name, "это кот")
	} else {
		dog3, okay := s.(*Dog)
		if okay {
			fmt.Println("Все таки", dog3.Name, " это собака!")
		}
	}

	var s2 Speaker = &Cat{Animal: Animal{Name: "Barsik"}, Color: "Gray"}
	switch typeAnimal := s2.(type) {
	case *Dog:
		fmt.Println(typeAnimal.Name, " Это собака!")
	case *Cat:
		fmt.Println(typeAnimal.Name, " Это кот!")
	default:
		fmt.Println("Непонятно что это за животное!")
	}
}

//v1.0
//работа с интерфейсами, встраивание структур, реализация методов интерфейса
//методы dog затеняют методы из родительской структуры
//кот юзает метод speak от animal т.к не реализован свой
//создан метод принимающий объект интерфейса, который выводит спик для каждой стркутуры свой

//v2.0
//можно создавать вложенные интерфейсы, создавая интерфейс, который содержит все методы вложенных
//из интерфейса можно доставать тип обратно и работать уже как с объектом определенного типа
//можно выяснить является ли определенным типом объект интерфейса через _,ok:=
//можно проверять какому типу принадлежит интерфейс через name.(type)
