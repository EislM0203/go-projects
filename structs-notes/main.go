package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type outputtable interface {
	Saver
	Display()
}

func main(){
	title, content := getNoteData()
	note, err := note.New(title, content)
	if err != nil {
		panic(err)
	}
	outputData(note)

	todoText := getUserInput("Text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		panic(err)
	}
	outputData(todo)
}

func outputData(data outputtable) {
	data.Display()
	saveData(data)
}

func saveData(data Saver) {
	err := data.Save()
	if err != nil {
		panic(err)
	}
	fmt.Println("Saving the data worked")
}

func getNoteData() (string, string) {
	title := getUserInput("Title:")
	content := getUserInput("Content:")
	return title, content
}

func getUserInput(infoText string) string {
	fmt.Printf("%v ", infoText)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}