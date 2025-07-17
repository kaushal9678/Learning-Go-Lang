package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)
type Note struct{
	Title string `json:"title"` // this is called struct tags
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
func (note Note) Display()  {
	fmt.Printf("\nYour note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}
func (note Note) Save() error  {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}
	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	},nil
}