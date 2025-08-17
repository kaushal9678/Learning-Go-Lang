package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)
type FileManager struct{
	InputFilePath string
	OutputFilePath string
}
func (fm FileManager)ReadLines() ([]string, error) {
	var prices []string
	file, error := os.Open(fm.InputFilePath)
	if error != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var price float64
		if _, err := fmt.Sscanf(scanner.Text(), "%f", &price); err == nil {
			prices = append(prices, fmt.Sprintf("%.2f", price))
		}else{
			return nil, errors.New("failed to read price from line")
		}
	}
	if err := scanner.Err(); err != nil {
        fmt.Println("Scanner error:", err)
		file.Close()
		return nil,err
    }
	
	return prices, nil
}

func (fm FileManager) WriteJSON(data interface{}) error{
	file, err := os.Create(fm.OutputFilePath)
	if err != nil{
		fmt.Println("Error creating file:", err)
		return errors.New(("failed to create file"));
	}
	//defer file.Close()
	if err := json.NewEncoder(file).Encode(data); err != nil {
		fmt.Println("Error writing JSON:", err)
		file.Close()
		return errors.New("failed to COVERT DATA TO JSON")
	}
	file.Close()
	return nil
}
func New(inputFilePath, outPutFiletPath string) *FileManager{
	return &FileManager{
		InputFilePath: inputFilePath,
		OutputFilePath: outPutFiletPath,
	}
}