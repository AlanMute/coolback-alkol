package repository

import (
	"fmt"

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

func (r *ModulePostgres) Delete(name string, courseName string) error {
	var course core.Course

	fmt.Printf("name of course = %s\n", courseName)

	if result := r.db.Where("name = ?", courseName).First(&course); result.Error != nil {
		return result.Error
	}

	fmt.Printf("course_id = %d name of course = %s\n", course.ID, name)

	var module core.Module
	if result := r.db.Where("course_id = ? AND name = ?", course.ID, name).First(&module); result.Error != nil {
		return result.Error
	}

	fmt.Printf("ID of deleted module = %d\n", module.ID)

	if result := r.db.Where("id = ?", module.ID).Unscoped().Delete(&module); result.Error != nil {
		return result.Error
	}

	return nil
}
