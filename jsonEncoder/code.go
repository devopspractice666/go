package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	Name  string  `json:"name"`
	Notes []*Note `json:"notes"`
}

type Data struct {
	Users []*User `json:"users"`
}

func (d *Data) String() string {
	if d == nil || len(d.Users) < 2 {
		return "Распарсилось, но второго пользователя нет!"
	}
	return "Распарсилось, пруф:\nИмя второго юзера:" + d.Users[1].Name
}
func main() {
	var info Data
	file, err := os.Open("data.json")
	if err != nil {
		panic("Ошибка с открытием файла")
	}

	//декодирование из файла в структуру
	decoder := json.NewDecoder(file)
	decoder.Decode(&info)
	fmt.Println(&info)
	for _, user := range info.Users {
		for _, note := range user.Notes {
			note.Title = "Стать хорошим devops инженером!"
		}
	}
	file.Close()

	//кодирование из структуры в файл
	file, err = os.OpenFile("storage.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic("Ошибка при открытии файла перезаписи")
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(info)
	file.Close()

	//парсинг якобы неизвестного json формата через мапу
	file, err = os.Open("shop.json")
	if err != nil {
		panic("Ошибка открытия файла shop.json")
	}
	data := make(map[string]any)
	secondDecoder := json.NewDecoder(file)
	secondDecoder.Decode(&data)
	for key, value := range data {
		fmt.Println("Ключ:", key, "Значение:", value)
	}
	file.Close()

	if name, ok := data["shop_name"]; ok {
		fmt.Println("Имя магазина:", name)
	}
}
