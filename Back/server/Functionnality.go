package OSINT

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pw := []byte(password)
	result, _ := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	return string(result)
}
func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

// func IsUserConnected(r *http.Request) bool {
// 	sessionCookie, errName := r.Cookie("session")
// 	if errName != nil {
// 		// If the cookie is not found, prompt the user to log in
// 		TplData.ProcessMessage = "Please log in"
// 		return false
// 	}
// 	// Verify the value of the "session" cookie
// 	userId := sessionCookie.Value
// 	if userId == "" {
// 		// If the cookie value is empty, prompt the user to log in
// 		TplData.ProcessMessage = "Please log in"
// 		return false
// 	}
// 	return true
// }
