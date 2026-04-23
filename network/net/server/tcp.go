package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {

	listener, err := net.Listen("tcp", ":8066")
	if err != nil {
		log.Fatal("Ошибка работы с портом! Невозможно слушать порт")

	}
	defer listener.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stop
		fmt.Println("\n Получен сигнал остановки, завершаем работу...")
		listener.Close()
	}()

	fmt.Println("Сервер запущен!")
	for {
		connect, err := listener.Accept()
		if err != nil {
			fmt.Println("Сервер не может устанавливать соединения")
			break
		}
		fmt.Println("Установка нового соединения!")
		wg.Add(1)
		go handle(connect)
	}
	fmt.Println("Ожидание завершения активных соединений...")

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Все соединения закрыты")
	case <-time.After(5 * time.Second):
		fmt.Println(" Таймаут ожидания, принудительное завершение")
	}

	fmt.Println("Сервер остановлен")

}

func handle(conn net.Conn) {
	defer conn.Close()
	defer wg.Done()
	conn.SetDeadline(time.Now().Add(5 * time.Minute))
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	writer.WriteString("установлено соединение\n")
	writer.Flush()
	for {
		conn.SetDeadline(time.Now().Add(5 * time.Minute))
		line, err := reader.ReadString('\n')
		if err != nil {

			fmt.Println("Закрыто соединение с клиентом")
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
