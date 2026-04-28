package note

import (
	"sort"
	"strconv"
)

type Note struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Date string `json:"date"`
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
