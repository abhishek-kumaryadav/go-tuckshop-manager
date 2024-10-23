package services

import (
	"log"
	"os"
)

func ReadReadmeFile(filePath string) (string, error) {
	data := ReadFile(filePath)
	return string(data), nil
}

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Critical error: %v", err)
	}
	return data
}
