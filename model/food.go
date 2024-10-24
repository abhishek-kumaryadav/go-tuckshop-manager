package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Food struct {
	ID    string
	Label string
	Price int
}

func ConvertToFoods(foodMap map[string]string) []Food {
	var foods []Food
	for k, v := range foodMap {
		start := strings.Index(k, "₹")
		if start == -1 {
			continue
		}
		price, err := strconv.Atoi(k[start+3:])
		if err != nil {
			// ... handle error
			panic(err)
		}
		end := strings.LastIndex(k[:start], " ")
		food := Food{v, k[:end], price}
		foods = append(foods, food)
	}
	return foods
}

func ConvertToString(foods []Food) string {
	var foodString strings.Builder
	for idx, food := range foods {
		foodString.WriteString(food.string(idx))
		foodString.WriteString("\n")
	}
	return foodString.String()
}

func (food *Food) string(idx int) string {
	var foodString strings.Builder
	foodString.WriteString(fmt.Sprintf("%d. %s %d -> %s", idx, food.Label, food.Price, food.ID))
	return foodString.String()
}
