package dao

// InboundBlouse model.
type InboundBlouse struct {
	ID             int    `json:"id"`
	BlouseID       int    `json:"blouse_id"`
	SKU            string `json:"sku"`
	Name           string `json:"name"`
	OrderAmount    int    `json:"order_amount"`
	ReceivedAmount int    `json:"received_amount"`
	PurchasePrice  int    `json:"purchase_price"`
	TotalPrice     int    `json:"total_price"`
	ReceiptNumber  string `json:"receipt_number"`
	Notes          string `json:"notes"`
	GenericModel
}
