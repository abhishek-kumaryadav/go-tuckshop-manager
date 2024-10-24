package repositories

import "go-tuckshop-manager/model"

func AddOrderTemplate(model.OrderTemplate) int {
	coll, _ := GetConnection()
	return 1
}
