package cmdline

import (
	"fmt"
	"go-tuckshop-manager/model"
	"go-tuckshop-manager/repositories"
	"go-tuckshop-manager/services"
)

func CreateOrderTemplateFlow() {
	fmt.Println(GetShop())
	input := services.ReadInputLineAsIntArray("Please enter space separated item IDs")
	foodMap := map[int]model.Food{}
	foodList := []model.Food{}
	for _, itemId := range input {
		foodList = append(foodList, foodMap[itemId])
	}
	orderTemplate := model.OrderTemplate{foodList}
	id := repositories.AddOrderTemplate(orderTemplate)
	fmt.Println(fmt.Sprintf("Added template order with ID: %d", id))
}
