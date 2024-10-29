package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error){
	if text == "" {
		return Todo{}, errors.New("text must not be empty strings")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display () {
	fmt.Printf("Your todo has the following text:\n%v", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonNote, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonNote, 0644)
}