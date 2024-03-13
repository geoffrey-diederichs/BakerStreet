package authentification

import (
	"github.com/golang-jwt/jwt"
	"time"
)
// Define a custom claims structure to include additional information in the token.
type CustomClaims struct {
    UserID string `json:"userId"`
    jwt.StandardClaims
}

// Function to generate a new JWT token.
func GenerateToken(userID string, secretKey []byte) (string, error) {
    // Set up our custom claims
    claims := CustomClaims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
            Issuer:    "your-app-name",
        },
    }

    // Create a new token object, specifying signing method and the claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with our secret
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}