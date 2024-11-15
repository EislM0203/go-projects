package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error){
	if title == "" || content == "" {
		return Note{}, errors.New("title and content must not be empty strings")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display () {
	fmt.Printf("Your note titled %v has the following content:\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonNote, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonNote, 0644)
}