package core

import (
	"time"

	"github.com/jinzhu/gorm"
)

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

type Sessions struct {
	gorm.Model
	RefreshToken   string
	ExpirationTime time.Time
}

type Email struct {
	gorm.Model

	Address string `gorm:"not null; unique"`
}
