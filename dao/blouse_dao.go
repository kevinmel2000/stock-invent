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
