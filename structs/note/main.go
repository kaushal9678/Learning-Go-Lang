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
type displayer interface {
	Display()
}
type outputtable interface {
	saver
	displayer // or you can directly add Display() method to the interface instead of adding displayer interface
}
func main() {
	printAnything(1)
	printAnything(123)
	printAnything(123.45)
	printAnything(true)
	printAnything([]string{"apple", "banana", "cherry"})
	printAnything(map[string]int{"apple": 1, "banana": 2, "cherry": 3})
	
	title, content := getNoteData()
	text := getToDoData()
	userNote, err := note.New(title, content)
	userTodo, err := todo.New(text)
	
	err =  displayData(userNote)
	if err != nil {
		return 
	}
	
	err = displayData(userTodo)
	if err != nil {
		return 
	}
	
}
func printAnything(data interface{}) {// here interface{} means you can pass any type of data to the function like object in Javascript} 
	// like in javascript we have typeof here we have switch
	switch data.(type) {
	case int:
		fmt.Printf("int %v\n", data)
	case string:
		fmt.Printf("int %v\n", data)
	case float64:
		fmt.Printf("float64 %v\n", data)
	case bool:
		fmt.Printf("bool %v\n", data)
	case []string:
		fmt.Printf("[]string %v\n", data)
	case map[string]int:
		fmt.Printf("map[string]int %v\n", data)
	default:
		fmt.Printf("unknown type %T\n", data)
	}

}
func printAnythingUsingSecondApproach(data interface{}) {
	intVal, ok := data.(int)
	if ok {
		fmt.Printf("int %v\n", intVal)
	}
	stringVal, ok := data.(string)
	if ok {
		fmt.Printf("string %v\n", stringVal)
	}
	floatVal, ok := data.(float64)
	if ok {
		fmt.Printf("float64 %v\n", floatVal)
	}
	boolVal, ok := data.(bool)
	if ok {
		fmt.Printf("bool %v\n", boolVal)
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
func displayData(data outputtable) error {
	data.Display()
	return saveData(data)
	
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