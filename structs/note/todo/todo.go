package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)
type Todo struct{
	Text string `json:"text"`
	
}
func (todo Todo) Display()  {
	fmt.Println(todo.Text)
}
func (todo Todo) Save() error  {
	fileName := strings.ReplaceAll(todo.Text, " ", "_")
	fileName = "todo.json" //strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
func New(text string) (Todo, error) {
	if text == ""{
		return Todo{}, errors.New("Invalid input")
	}
	return Todo{
		Text: text,
	},nil
}