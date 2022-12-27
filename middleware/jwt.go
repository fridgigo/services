package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(email string) (string, error) {

	// Secret key used to sign the token
	secret := []byte("my-secret-key")

	// Claims to be stored in the token
	claims := jwt.MapClaims{
		"iss": "my-issuer",
		"sub": "my-subject",
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]string{
			"email": email,
		},
	}

	// Create and sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		// Error handling
		fmt.Println(err)
		return "", err
	}
	return signedToken, nil
}

func VerifyJWT() (string, error) {
	return "", nil
}
