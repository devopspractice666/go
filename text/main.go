package main

import (
	"encoding/json"
	"fmt"
	. "myapp/note"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Serv struct {
	list []*Note
}

func NewServer() *Serv {
	data := GetFromJson()
	list := data.Notes
	return &Serv{list: list}
}

func (l Serv) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	listNote := GetFromJson()

	b, err := json.MarshalIndent(listNote.Notes, "", " ")
	if err != nil {
		fmt.Println("Ошибка кодирования в json формат")
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println("Ошибка записи в ответ json формата")
	}
}

func (l Serv) GetNoteByText(w http.ResponseWriter, r *http.Request) {
	listNote := GetNoteByText(mux.Vars(r)["text"])
	if len(listNote) == 0 {
		w.WriteHeader(404)
		return
	}
	b, err := json.MarshalIndent(listNote, "", " ")
	if err != nil {
		fmt.Println("Ошибка кодирования в json формат")
	}

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("Ошибка записи в ответ json формата")
	}
}

func (l Serv) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(400)
		return
	}
	note := GetNoteByID(id)
	if note == nil {
		w.WriteHeader(404)
		return
	}
	b, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(b)
}

func (l Serv) AddNoteToDatabase(w http.ResponseWriter, r *http.Request) {
	var data map[string]any
	json.NewDecoder(r.Body).Decode(&data)
	if _, ok := data["text"]; !ok {
		w.WriteHeader(400)
		return
	}
	if len(data) != 1 {
		w.WriteHeader(400)
		return
	}
	text := data["text"].(string)
	AddNoteByText(text)
	message := "Данные добавлены в базу успешно"
	b, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(201)
		w.Write(b)
	}
}

func (l Serv) RemoveNoteFromDatabase(w http.ResponseWriter, r *http.Request) {
	var data map[string]any
	json.NewDecoder(r.Body).Decode(&data)
	if _, ok := data["text"]; !ok {
		w.WriteHeader(400)
		return
	}
	if len(data) != 1 {
		w.WriteHeader(400)
		return
	}
	text := data["text"].(string)
	RemoveNotesByText(text)
	message := "Данные удалены из базы успешно (если они есть)"
	b, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(204)
		w.Write(b)
	}
}

func (l Serv) ChangeText(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	old := vars["p1"]
	new := vars["p2"]
	ChangeNoteText(old, new)
	w.WriteHeader(204)

}

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
