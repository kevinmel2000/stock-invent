package dao

// Blouse object with generic_data included.
type Blouse struct {
	ID    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	GenericModel
}
