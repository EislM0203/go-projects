package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note/note"
)

func main(){
	title, content := getNoteData()
	note, err := note.New(title, content)

	if err != nil {
		panic(err)
	}

	note.Display()
	err = note.Save()
	if err != nil {
		panic(err)
	}
	fmt.Println("Saving the note worked")
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