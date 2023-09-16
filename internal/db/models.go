package db

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	NameFolder  string
}

type Module struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	CoursesID   int
	NameFolder  string
}

type Lesson struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	ModulesID   int
	NameFolder  string
}
