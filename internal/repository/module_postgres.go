package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type ModulePostgres struct {
	db *gorm.DB
}

func NewModulePostgres(db *gorm.DB) *ModulePostgres {
	return &ModulePostgres{db: db}
}

func (r *ModulePostgres) Add(name string, description string, courseName string, folderName string) error {
	newModule := core.Module{
		Name:        name,
		Description: description,
		NameFolder:  folderName,
	}

	if result := r.db.Create(&newModule); result.Error != nil {
		return result.Error
	}

	var course core.Course
	if result := r.db.Where("name = ?", courseName).First(&course); result.Error != nil {
		return result.Error
	}

	course.Modules = append(course.Modules, newModule)
	if result := r.db.Save(&course); result.Error != nil {
		return result.Error
	}

	return nil
}
