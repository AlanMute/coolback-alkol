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
	var course core.Course
	if result := r.db.Where("name = ?", courseName).First(&course); result.Error != nil {
		return result.Error
	}

	newModule := core.Module{
		Name:        name,
		Description: description,
		NameFolder:  folderName,
		CourseID:    course.ID,
	}

	if result := r.db.Create(&newModule); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ModulePostgres) Delete(id uint) (string, error) {
	var module core.Module
	if result := r.db.Where("id = ?", id).First(&module); result.Error != nil {
		return "", result.Error
	}

	if result := r.db.Where("id = ?", id).Unscoped().Delete(&module); result.Error != nil {
		return "", result.Error
	}

	return module.NameFolder, nil
}
