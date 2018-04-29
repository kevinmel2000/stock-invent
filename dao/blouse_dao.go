package dao

import (
	"context"
)

type BlouseDao struct{}

func NewBlouseDao(ctx context.Context) *BlouseDao {
	return &BlouseDao{}
}

// FindAll return all array of blouse[].
func (blouseDao BlouseDao) FindAll(ctx context.Context) ([]Blouse, error) {
	db := initDB()
	defer db.Close()

	var blouse []Blouse
	err := db.Find(&blouse).Error

	return blouse, err
}
