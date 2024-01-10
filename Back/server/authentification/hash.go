package authentification

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	// Generate a hashed password
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}
