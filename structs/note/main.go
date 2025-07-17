package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)
type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	text := getToDoData()
	userNote, err := note.New(title, content)
	userTodo, err := todo.New(text)
	
	userNote.Display()

	userTodo.Display()

	err = saveData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = saveData(userTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
	
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Note saved successfully")
	return nil
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}
func getToDoData() string {
	text := getUserInput("ToDo  text:")
	return text
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}