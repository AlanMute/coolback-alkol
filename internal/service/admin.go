package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type AdminService struct {
	repo repository.Course
}

func NewAdminService(repo repository.Course) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) SignIn(login string, pass string) (core.Tokens, error) {

	return core.Tokens{}, nil
}
