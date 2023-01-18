package models

type Item struct {
	Item_id     int    `json:"item_id" gorm:"primary_key"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    *int   `json:"order_id"`
	Order       Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ItemRequestDTO struct {
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
