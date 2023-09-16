package repository

import "github.com/jinzhu/gorm"

type LessonPostgres struct {
	db *gorm.DB
}

func NewLessonPostgres(db *gorm.DB) *LessonPostgres {
	return &LessonPostgres{db: db}
}

func (r *LessonPostgres) Add(name string, description string, id int, fileName string) error {

	return nil
}

func (r *LessonPostgres) Get(name string) error {

	return nil
}
