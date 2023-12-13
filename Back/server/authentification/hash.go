package authentification

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	fmt.Println(password)
	// Generate a hashed password
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}
