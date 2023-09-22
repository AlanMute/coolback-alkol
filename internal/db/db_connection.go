package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/KrizzMU/coolback-alkol/internal/config"
	"github.com/KrizzMU/coolback-alkol/internal/core"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres", config.GetConnectionString())
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&core.Course{})
	db.AutoMigrate(&core.Module{})
	db.AutoMigrate(&core.Course{})

	return db
}
