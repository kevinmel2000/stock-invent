package dao

import "context"

// InboundBlouseDao object.
type InboundBlouseDao struct{}

// NewInboundBlouseDao used to access member of blouse_dao object.
func NewInboundBlouseDao(ctx context.Context) *InboundBlouseDao {
	return &InboundBlouseDao{}
}

// Insert will create a record from Blouse object.
func (inboundBlouseDao InboundBlouseDao) Insert(ctx context.Context, inboundBlouse InboundBlouse) error {
	db := InitDB()
	defer db.Close()

	err := db.Create(&inboundBlouse).Error

	return err
}
