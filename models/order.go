package models

import "time"

type Order struct {
	Order_id      int       `json:"order_id" gorm:"primary_key"`
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	Items         []Item    `gorm:"foreignKey:Order_id;constraint:OnDelete:CASCADE"`
}

type OrderRequestDTO struct {
	Ordered_at    time.Time        `json:"orderedAt"`
	Customer_name string           `json:"customerName"`
	Items         []ItemRequestDTO `json:"items"`
}
