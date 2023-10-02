package core

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model

	Name        string `gorm:"not null; unique"`
	Description string `gorm:"not null"`
}

type Module struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	OrderID     uint   `gorm:"not null"`
	CourseID    uint

	Course Course `gorm:"foreignKey:CourseID" json:"-"`
}

type Lesson struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	OrderID     uint   `gorm:"not null"`
	ModuleID    uint

	Module Module `gorm:"foreignKey:ModuleID" json:"-"`
}

type User struct {
	gorm.Model

	Login    string `gorm:"not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"not null"`
}

type Email struct {
	gorm.Model

	Email string `gorm:"not null"`
}
