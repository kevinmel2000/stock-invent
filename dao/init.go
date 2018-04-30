package dao

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// InitDB used to create a db when one is not created.
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./clothes_invent.db")
	db.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}

	if !db.HasTable(&Blouse{}) {
		db.CreateTable(&Blouse{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Blouse{})
	}

	if !db.HasTable(&InboundBlouse{}) {
		db.CreateTable(&InboundBlouse{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&InboundBlouse{})
	}

	return db
}
