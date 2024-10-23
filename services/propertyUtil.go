package services

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

var envMap = map[string]string{}
var ENV_FILE_NAME = "./.env"

func InitEnvProperties() map[string]string {
	data := ReadFile(ENV_FILE_NAME)
	if len(data) == 0 {
		return envMap
	}
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineStr := string(line)
		keyVal := strings.Split(lineStr, "=")
		envMap[keyVal[0]] = keyVal[1]
	}
	return envMap
}

func GetEnvProperties() map[string]string {
	return envMap
}

func UpdateEnvFile(newMap map[string]string) {
	var fileString strings.Builder
	for k, v := range newMap {
		fileString.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}

	file, err := os.Create(ENV_FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(fileString.String()); err != nil {
		log.Printf("Critical error: %v", err)
	}
}
