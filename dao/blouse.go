package dao

import "time"

// Blouse object with generic_data included.
type Blouse struct {
	ID    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	GenericModel
}

// GenericModel object.
type GenericModel struct {
	CreatedAt time.Time `gorm:"null" json:"created_at"`
	CreatedBy string    `gorm:"null" json:"created_by"`
	UpdatedAt time.Time `gorm:"null" json:"updated_at"`
	UpdatedBy string    `gorm:"null" json:"updated_by"`
}
