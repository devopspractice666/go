package note

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Note struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Date string `json:"date"`
}
type NoteList struct {
	NextId int     `json:"nextid"`
	Notes  []*Note `json:"notes"`
}

func (note *Note) NoteString() string {
	return "NoteID: " + strconv.Itoa(note.Id) + "<pad>Text: " + note.Text + "<pad>" + "CreateDate:" + note.Date + "\n"
}

func SortById(notes []*Note) []*Note {
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Id < notes[j].Id
	})
	return notes
}

func AddToJson(notes NoteList) error {
	jsondata, err2 := json.MarshalIndent(&notes, "", "  ")
	if err2 != nil {
		return err2
	}
	err2 = os.WriteFile("Notes.json", jsondata, 0644)
	if err2 != nil {
		return err2
	}
	return nil

}

func GetFromJson() *NoteList {
	data, _ := os.ReadFile("Notes.json")
	var listNote NoteList
	err := json.Unmarshal(data, &listNote)
	if err != nil {
		return nil
	}
	return &listNote
}

func AddToList(note *Note) error {
	listNote := GetFromJson()
	note.Id = listNote.NextId
	listNote.NextId++
	listNote.Notes = append(listNote.Notes, note)
	AddToJson(*listNote)
	return nil
}

func AddNoteByText(text string) error {
	note := &Note{Text: text, Date: time.Now().Format("02-01-2006 15:04:05")}
	AddToList(note)
	return nil
}

func FindNoteByText(text string) string {
	var isfind bool = false
	listNote := GetFromJson()
	var slice []*Note
	for _, note := range listNote.Notes {
		if strings.Contains(strings.ToLower(note.Text), strings.ToLower(text)) {
			isfind = true
			slice = append(slice, note)
		}
	}
	var answer string
	for _, note := range slice {
		answer += "Найдена заметка с id:" + strconv.Itoa(note.Id) + " Дата создания: " + note.Date + "\n"
	}
	if isfind {
		return answer
	}
	return "Заметка с вашим текстом не найдена"
}

func GetNoteByText(text string) []*Note {
	listNote := GetFromJson()
	var slice []*Note
	for _, note := range listNote.Notes {
		if note.Text == text {
			slice = append(slice, note)
		}
	}
	return slice
}
func GetNoteByID(id int) *Note {
	listNote := GetFromJson()
	for _, note := range listNote.Notes {
		if note.Id == id {
			return note
		}
	}
	return nil
}
func RemoveNotesByText(text string) {
	listNote := GetFromJson()
	var slice []*Note
	for _, note := range listNote.Notes {
		if note.Text != text {
			slice = append(slice, note)
		}
	}
	listNote.Notes = slice
	AddToJson(*listNote)
}

func ChangeNoteText(text string, changedText string) {
	listNote := GetFromJson()
	for _, note := range listNote.Notes {
		if note.Text == text {
			note.Text = changedText
		}
	}
	AddToJson(*listNote)

}

type WordCount struct {
	word  string
	count int
}

func TopPopularWords(n int) {
	listNote := GetFromJson()
	count := make(map[string]int)
	for _, note := range listNote.Notes {
		words := strings.Split(note.Text, " ")
		for _, word := range words {
			count[strings.ToLower(word)]++
		}
	}
	var slice []WordCount
	for key, value := range count {
		slice = append(slice, WordCount{word: strings.ToLower(key), count: value})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].count > slice[j].count
	})
	var top int
	if len(slice) < n {
		top = len(slice)
	} else {
		top = n
	}
	fmt.Println("Топ ", top, " самых популярных слов в заметках:")
	for i := 0; i < top; i++ {
		fmt.Println(slice[i].word, " Встретилось ", slice[i].count, " раз")
	}

}
