package cmdline

import (
	"go-tuckshop-manager/model"
	"go-tuckshop-manager/repositories"
	"go-tuckshop-manager/services"
)

func GetShop() string { //todo get from mongo
	foodMap := services.GetFoodMap()
	foods := model.ConvertToFoods(foodMap)
	return model.ConvertToString(foods)
}

func UpdateShop() {
	foodMap := services.GetFoodMap()
	foods := model.ConvertToFoods(foodMap)
	repositories.UpdateShopDatabase(foods)
}
