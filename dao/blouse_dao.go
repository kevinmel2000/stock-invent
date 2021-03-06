package dao

import (
	"context"
)

// BlouseDao object.
type BlouseDao struct{}

// NewBlouseDao used to access member of blouse_dao object.
func NewBlouseDao(ctx context.Context) *BlouseDao {
	return &BlouseDao{}
}

// FindAll return all array of blouse[].
func (blouseDao BlouseDao) FindAll(ctx context.Context) ([]Blouse, error) {
	db := InitDB()
	defer db.Close()

	var blouse []Blouse
	err := db.Find(&blouse).Error

	return blouse, err
}

// FindOne return one record of blouse by id
func (blouseDao BlouseDao) FindOne(ctx context.Context, id int) (Blouse, error) {
	db := InitDB()
	defer db.Close()

	var blouse Blouse
	err := db.Find(&blouse, id).Error

	return blouse, err
}

// Insert will create a record from Blouse object.
func (blouseDao BlouseDao) Insert(ctx context.Context, blouse Blouse) error {
	db := InitDB()
	defer db.Close()

	err := db.Create(&blouse).Error

	return err
}

// Update will edit one record of blouse by id
func (blouseDao BlouseDao) Update(ctx context.Context, id int, newblouse Blouse) error {
	db := InitDB()
	defer db.Close()

	var blouse Blouse
	err := db.First(&blouse, id).Error
	if err != nil {
		return err
	}

	err = db.Save(newblouse).Error

	return err
}

// Delete will delete one record of blouse by id
func (blouseDao BlouseDao) Delete(ctx context.Context, id int) error {
	db := InitDB()
	defer db.Close()

	var blouse Blouse
	err := db.First(&blouse, id).Error
	if err != nil {
		return err
	}

	err = db.Delete(blouse).Error

	return err
}
