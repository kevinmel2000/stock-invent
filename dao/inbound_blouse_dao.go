package dao

import "context"

// InboundBlouseDao object.
type InboundBlouseDao struct{}

// NewInboundBlouseDao used to access member of blouse_dao object.
func NewInboundBlouseDao(ctx context.Context) *InboundBlouseDao {
	return &InboundBlouseDao{}
}

// Insert will create a record from object.
func (inboundBlouseDao InboundBlouseDao) Insert(ctx context.Context, inboundBlouse InboundBlouse) error {
	db := InitDB()
	defer db.Close()

	err := db.Create(&inboundBlouse).Error

	return err
}

// FindOne return one record by id.
func (inboundBlouseDao InboundBlouseDao) FindOne(ctx context.Context, id int) (InboundBlouse, error) {
	db := InitDB()
	defer db.Close()

	var inboundBlouse InboundBlouse
	err := db.Find(&inboundBlouse, id).Error

	return inboundBlouse, err
}

// FindAll return all record.
func (inboundBlouseDao InboundBlouseDao) FindAll(ctx context.Context) ([]InboundBlouse, error) {
	db := InitDB()
	defer db.Close()

	var inboundBlouse []InboundBlouse
	err := db.Find(&inboundBlouse).Error

	return inboundBlouse, err
}

// Update update single record by id.
func (inboundBlouseDao InboundBlouseDao) Update(ctx context.Context, id int, NewInboundBlouse InboundBlouse) error {
	db := InitDB()
	defer db.Close()

	var inboundBlouse InboundBlouse
	err := db.First(&inboundBlouse, id).Error
	if err != nil {
		return err
	}

	err = db.Save(NewInboundBlouse).Error

	return err
}

// Delete will delete one record of blouse by id
func (inboundBlouseDao InboundBlouseDao) Delete(ctx context.Context, id int) error {
	db := InitDB()
	defer db.Close()

	var inboundBlouse InboundBlouse
	err := db.First(&inboundBlouse, id).Error
	if err != nil {
		return err
	}

	err = db.Delete(inboundBlouse).Error

	return err
}
