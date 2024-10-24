package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputLine(input string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(input)

	key, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return key
}

func ReadInputLineAsIntArray(input string) []int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(input)

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	lineItem := strings.Split(line, " ")
	parsedItem := []int{}
	for _, item := range lineItem {
		itemInt, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		parsedItem = append(parsedItem, itemInt)
	}
	return parsedItem
}
