package model

import "time"

type Data struct {
	CreatedAt time.Time `gorm:"null" json:"created_at"`
	CreatedBy string    `gorm:"null" json:"created_by"`
	UpdatedAt time.Time `gorm:"null" json:"updated_at"`
	UpdatedBy string    `gorm:"null" json:"updated_by"`
}

func NewData() Data {
	return Data{
		CreatedAt: time.Now(),
		CreatedBy: "",
	}
}
