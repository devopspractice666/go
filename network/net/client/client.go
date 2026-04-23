package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8066")
	if err != nil {
		panic("Ошибка подключения к серверу! Сервер недоступен")
	}
	defer conn.Close()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			connfromserv, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Сервер разорвал соединение")
				return
			}
			fmt.Print("От сервера:", connfromserv)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text += "\n"
		if text == "exit\n" {
			conn.Write([]byte("exit\n"))
			fmt.Println("Завершение работы клиента...")
			return
		}
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Соединение с сервером нарушено!")
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
	}
}
