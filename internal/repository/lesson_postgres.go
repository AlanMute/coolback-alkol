package repository

import (
	"fmt"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type LessonPostgres struct {
	db *gorm.DB
}

func NewLessonPostgres(db *gorm.DB) *LessonPostgres {
	return &LessonPostgres{db: db}
}

func (r *LessonPostgres) Add(name string, description string, fileName string, courseName string, moduleName string) error {
	var course core.Course
	if result := r.db.Where("name = ?", courseName).First(&course); result.Error != nil {
		return result.Error
	}

	var module core.Module
	if result := r.db.Where("name = ? AND course_id = ?", moduleName, course.ID).First(&module); result.Error != nil {
		return result.Error
	}

	newLesson := core.Lesson{
		Name:        name,
		Description: description,
		NameFile:    fileName,
		ModuleID:    module.ID,
	}

	if result := r.db.Create(&newLesson); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *LessonPostgres) Delete(id uint) (string, error) {
	var lesson core.Lesson
	if result := r.db.Where("id = ?", id).First(&lesson).Unscoped().Delete(&lesson); result.Error != nil {
		return "", result.Error
	}

	if result := r.db.Where("id = ?", id).Unscoped().Delete(&lesson); result.Error != nil {
		return "", result.Error
	}

	fmt.Printf("NameFile = %s", lesson.NameFile)

	return lesson.NameFile, nil
}

func (r *LessonPostgres) Get(name string) error {

	return nil
}
