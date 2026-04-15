package main

import (
	"encoding/json"
	"fmt"
	"os"
	. "parse/models"
)

func main() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла ")
		panic("ОШибка при работе с файлом")
	}
	var fullInfo Data
	err = json.Unmarshal(data, &fullInfo)
	fmt.Println("Номер телефона ресторана:", fullInfo.Restaurant.Phone)

}
