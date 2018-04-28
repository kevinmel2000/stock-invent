package model

import "context"

// Item object with generic_data included.
type Item struct {
	ID    int    `json:"id"`
	Sku   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Data
}

// ItemDataModel object.
type ItemDataModel struct{}

// NewItemModel return object of ItemDataModel.
func NewItemModel(ctx context.Context) *ItemDataModel {
	return &ItemDataModel{}
}

// GetMany return all array of items[].
func (itemDataModel ItemDataModel) GetMany(ctx context.Context) ([]Item, error) {
	db := initDB()
	defer db.Close()

	var items []Item
	err := db.Find(&items).Error

	return items, err
}
