package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Log struct {
	Time    string
	Name    string
	Level   string
	Message string
}

func NewLog(time, name, level, message string) *Log {
	return &Log{
		Time:    time,
		Name:    name,
		Level:   level,
		Message: message,
	}
}
func PrintAllLogsWholeInfo(logs []*Log) {
	fmt.Println("Info about all logs from input line")
	fmt.Println(strings.Repeat("-", 50))
	for _, info := range logs {
		fmt.Println("ServerName:", info.Name, " Time:", info.Time, " Level:", info.Level, " message:", info.Message)
	}
}
func PrintErrors(logs []*Log) {
	fmt.Println("ERROR LOGS:")
	fmt.Println(strings.Repeat("-", 100))
	for _, log := range logs {
		if log.Level == "ERROR" {
			fmt.Printf("[ %-14s] %-10s %s\n", log.Time, log.Name, log.Message) //для ровных таблиц в выводе
		}
	}
}
func ParseInputLogs(input string) []*Log {

	exactservers := strings.Split(input, ";")
	var trimexact string
	var logs []*Log
	var time, name, level, message string
	for _, info := range exactservers {
		trimexact = strings.TrimSpace(info)
		internal := strings.SplitN(trimexact, " ", 4)
		for _, test := range internal {
			lastbeat := strings.Split(test, ":")
			switch lastbeat[0] {
			case "timestamp":
				time = lastbeat[1]
			case "server":
				name = lastbeat[1]
			case "level":
				level = lastbeat[1]
			case "message":
				message = lastbeat[1]
			}
		}
		log := NewLog(time, name, level, message)
		logs = append(logs, log)
		time, name, level, message = "", "", "", ""
	}
	return logs
}
func main() {
	var data string
	var file *os.File
	if len(os.Args) <= 2 {
		if len(os.Args) == 1 {
			file, _ = os.Open("testdata/logs.txt")
		} else {
			file, _ = os.Open(os.Args[1])
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			data = data + line
		}
		PrintErrors(ParseInputLogs(data))
	} else {
		fmt.Println("Введи 1 имя файла с логами, для вывода логов ошибок, как параметр (скрипт принимает ровно 1 параметр!)")
	}
}
