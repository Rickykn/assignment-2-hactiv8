package controllers

import (
	"fmt"

	"github.com/Rickykn/assignment-2-hactiv8.git/database"
	"github.com/Rickykn/assignment-2-hactiv8.git/models"
	"gorm.io/gorm"
)

func CreateOrder(orderInput *models.OrderRequestDTO) *models.Order {
	db := database.Get()
	var items []models.Item

	newOrder := &models.Order{
		Customer_name: orderInput.Customer_name,
		Ordered_at:    orderInput.Ordered_at,
	}

	db.Create(newOrder)

	for _, val := range orderInput.Items {
		newItem := &models.Item{
			Item_code:   val.Item_code,
			Description: val.Description,
			Quantity:    val.Quantity,
			Order_id:    &newOrder.Order_id,
		}
		items = append(items, *newItem)
		db.Create(items)
	}

	newOrder.Items = items

	return newOrder
}

func GetOrder() []*models.Order {
	db := database.Get()
	var orders []*models.Order

	db.Preload("Items", func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Order")
	}).Find(&orders)

	return orders

}

func DeleteOrder(id int) int64 {
	db := database.Get()
	fmt.Println(id)
	result := db.Delete(&models.Order{}, id)
	return result.RowsAffected
}
