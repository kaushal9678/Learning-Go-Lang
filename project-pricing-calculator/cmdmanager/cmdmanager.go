package cmdmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

type CommandManager struct {

}

func (fm CommandManager)ReadLines() ([]string, error) {
	var prices[]string;

	for {
		var input string
		fmt.Println("Please enter a price and press Enter (or type 'exit' and press Enter to finish):")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}
		if input == "" {
		fmt.Println("No value entered, please enter a price or 'exit'.")
		continue
	}
		  prices = append(prices, input)
	}

	if len(prices) == 0 {
		return nil, fmt.Errorf("no prices were entered")
	}
	return prices, nil
}
func (fm CommandManager) WriteJSON(data interface{}) error{
	fileName := "commandArguments.json"
	jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return fmt.Errorf("error marshaling JSON: %v", err)
    }
    file, err := os.Create(fileName)
    if err != nil {
        return fmt.Errorf("error creating file: %v", err)
    }
    defer file.Close()
    _, err = file.Write(jsonData)
    if err != nil {
        return fmt.Errorf("error writing to file: %v", err)
    }
    fmt.Printf("JSON data written to %s\n", fileName)
    return nil

}

func New() CommandManager{
	return  CommandManager{}
}