package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB is the connection string to the database
	db *gorm.DB
)

// SetupDatabaseConnection creates a connection to the database
func Connect() {
	database, err := gorm.Open("mysql", "root:password@tcp(@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!")
	}
	db = database

}

// GetDB returns a connection to the database
func GetDB() *gorm.DB {
	return db
}
