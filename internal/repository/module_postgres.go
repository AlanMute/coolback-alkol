package repository

import "github.com/jinzhu/gorm"

type ModulePostgres struct {
	db *gorm.DB
}

func NewModulePostgres(db *gorm.DB) *ModulePostgres {
	return &ModulePostgres{db: db}
}

func (r *ModulePostgres) Add(name string, description string, id int, folderName string) error {

	return nil
}
