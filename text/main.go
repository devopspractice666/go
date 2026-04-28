package main

import (
	"fmt"
	. "myapp/server"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ser := NewServer()
	router := mux.NewRouter()

	router.Path("/notes").Methods("GET").HandlerFunc(ser.GetAllNotes)
	router.Path("/notes/search/{text}").Methods("GET").HandlerFunc(ser.GetNoteByText)
	router.Path("/notes").Methods("POST").HandlerFunc(ser.AddNoteToDatabase)
	router.Path("/notes").Methods("DELETE").HandlerFunc(ser.RemoveNoteFromDatabase)
	router.Path("/notes").Methods("PATCH").Queries("oldText", "{p1}", "newText", "{p2}").HandlerFunc(ser.ChangeText)
	router.Path("/notes/{id}").Methods("GET").HandlerFunc(ser.GetNoteByID)

	err := http.ListenAndServe(":9066", router)
	if err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
}
