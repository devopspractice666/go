package main

import (
	"fmt"
	"strings"
)

type Server struct {
	Name   string
	Cpu    string
	Memory string
	Status string
}

func NewServer(name, cpu, memory, status string) *Server {
	return &Server{
		Name:   name,
		Cpu:    cpu,
		Memory: memory,
		Status: status,
	}
}

func (r *Server) ShowServer() {
	var status string
	if r.Status == "running" {
		status = "Active"
	} else {
		status = "Inactive"
	}

	fmt.Println("Сервак:", r.Name, "Находится в состоянии:", status, "Кол-во процессоров и памяти: ", r.Cpu, ",", r.Memory)
}

func listActive(all []*Server) {
	fmt.Println("Список активных серверов:")
	for _, serv := range all {
		if serv.Status == "running" {
			fmt.Println("Сервер:", serv.Name)
		}
	}
}

func parsing(inp string) []*Server {
	var slice []string = strings.Split(inp, ";")
	var name, cpu, memory string
	var status string
	var servers []*Server
	for _, str := range slice {
		withoutspace := strings.TrimSpace(str)
		//fmt.Println(withoutspace)
		internalSlice := strings.Split(withoutspace, " ")
		for _, str2 := range internalSlice {
			fields := strings.Split(str2, ":")
			switch fields[0] {
			case "server":
				name = fields[1]
			case "status":
				status = fields[1]
			case "cpu":
				cpu = fields[1]
			case "mem":
				memory = fields[1]
			}

		}
		servers = append(servers, NewServer(name, cpu, memory, status))
	}
	name = ""
	cpu = ""
	memory = ""
	status = ""

	return servers
}

func main() {
	input := "server:web01 status:running cpu:4 mem:8; server:db01 status:stopped cpu:8 mem:16; server:cache01 status:running cpu:2 mem:4"
	listActive(parsing(input))
	input4 := "server:web01 status:running cpu:4 mem:8; server:db01 status:running cpu:8 mem:16; server:cache01 status:running cpu:2 mem:4"
	listActive(parsing(input4))
}
