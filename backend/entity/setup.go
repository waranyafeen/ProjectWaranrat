package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("Projectwaranrat.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	autoMigrate(database)
	db = database
}

// Migrate the schema
func autoMigrate(database *gorm.DB) {
	database.AutoMigrate(
		&Employee{},
		&Precede{},
		&Position{},
		&Gender{},
		&User{},
		&Role{},
		&Ticket{},
		&Departure{},
		
	)
}
