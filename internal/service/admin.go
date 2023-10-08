package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
)

type AdminService struct {
	tockenManager auth.TokenManager
	repo          repository.Session
}

func NewAdminService(t auth.TokenManager, r repository.Session) *AdminService {
	return &AdminService{
		tockenManager: t,
		repo:          r,
	}
}

func (s *AdminService) SignIn(login string, pass string) (core.Tokens, error) {
	if login != os.Getenv("ADM_LOGIN") || pass != os.Getenv("ADM_PASSWORD") {
		return core.Tokens{}, errors.New("Uncorrect login or password!")
	}

	accessToken, err := s.tockenManager.NewAccessToken("admin", time.Second)

	if err != nil {
		return core.Tokens{}, err
	}

	refreshToken, err := s.tockenManager.NewRefreshToken()

	if err != nil {
		return core.Tokens{}, err
	}

	t := time.Now().AddDate(0, 0, 30)

	session := core.Sessions{
		RefreshToken:   refreshToken,
		ExpirationTime: t,
	}

	fmt.Println(t, refreshToken)
	if err := s.repo.Add(session); err != nil {

		return core.Tokens{}, nil
	}

	token := core.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
