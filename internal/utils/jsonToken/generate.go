package jsontoken

import (
	"github.com/All-Khwarizmi/golang-blog/internal/models"
	"github.com/All-Khwarizmi/golang-blog/internal/utils"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(email string, name string) (string, error) {
	// Set as env var
	secret := []byte(utils.GoDotEnvVariable("SECRET_KEY"))
	// Create the Claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		Email:            email,
		Name:             name,
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
