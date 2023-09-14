package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/KrizzMU/coolback-alkol/config"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres", config.GetConnectionString())
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Module{})
	db.AutoMigrate(&Lesson{})

	return db
}
