package db

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/KrizzMU/coolback-alkol/internal/core"
)

func GetDefaultAdmin() (core.User, error) {
	var admin core.User

	hash := sha512.New()

	password := "admin"

	_, err := hash.Write([]byte(password))
	if err != nil {
		return core.User{}, err
	}

	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	admin.Login = "admin"
	admin.Password = hashedPassword
	admin.IsAdmin = true

	return admin, nil
}
