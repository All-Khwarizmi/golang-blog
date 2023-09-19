package models

import "github.com/golang-jwt/jwt/v4"

// Data that will be in the token
type UserClaim struct {
	jwt.RegisteredClaims
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type User struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email"  binding:"required"`
}
