package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type SessionPostgres struct {
	db *gorm.DB
}

func NewSessionPostgres(db *gorm.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (r *SessionPostgres) Add(session core.Sessions) error {
	if result := r.db.Create(&session); result.Error != nil {
		return result.Error
	}

	return nil
}
