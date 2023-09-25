package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type CoursePostgres struct {
	db *gorm.DB
}

func NewCoursePostgres(db *gorm.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (r *CoursePostgres) Add(name string, description string, folderName string) error {
	newCourse := core.Course{
		Name:        name,
		Description: description,
		NameFolder:  folderName,
	}

	if result := r.db.Create(&newCourse); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CoursePostgres) Delete(name string) error {
	var course core.Course

	if result := r.db.Where("name = ?", name).First(&course); result.Error != nil {
		return result.Error
	}

	if result := r.db.Where("name = ?", name).Unscoped().Delete(&course); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CoursePostgres) GetByName(name string) ([]core.Course, error) {

	var courses []core.Course
	r.db.Where("name ILIKE ?", "%"+name+"%").Find(&courses)

	return courses, nil
}

func (r *CoursePostgres) GetAll() ([]core.Course, error) {

	var courses []core.Course
	r.db.Find(&courses)
	return courses, nil
}
