package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8066")
	if err != nil {
		log.Fatal("Ошибка работы с портом! Невозможно слушать порт")

	}
	fmt.Println("Сервер запущен!")
	for {
		connect, err := listener.Accept()
		if err != nil {
			log.Fatal("Ошибка при установлении tcp соединения!")
			panic("Ошибка сервера")
		}
		fmt.Println("Установка нового соединения!")
		go handle(connect)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	writer.WriteString("установлено соединение\n")
	writer.Flush()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("connection lost")
			fmt.Println("client:", line)
			return
		}
		if line == "exit\n" {
			conn.Write([]byte("Сервер закрывает соединение\n"))
			fmt.Println("Закрытие соединения с сервером")
			return
		}
		fmt.Print(line)
		line = "ECHO " + line
		conn.Write([]byte(line))
	}

}
