package dao

import "time"

// GenericModel object.
type GenericModel struct {
	CreatedAt time.Time `gorm:"null" json:"created_at"`
	CreatedBy string    `gorm:"null" json:"created_by"`
	UpdatedAt time.Time `gorm:"null" json:"updated_at"`
	UpdatedBy string    `gorm:"null" json:"updated_by"`
}
