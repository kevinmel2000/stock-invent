package model

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type ModelModule struct{}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./inventory.db")
	db.LogMode(true)
	if err != nil {
		fmt.Errorf("Could not open db: %v", err)
	}

	if !db.HasTable(&Item{}) {
		db.CreateTable(&Item{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Item{})
	}

	return db
}

func initSQL() *sql.DB {
	db, _ := sql.Open("sqlite3", "./inventory.db")
	return db
}
