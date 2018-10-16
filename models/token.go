package models

import jwt "github.com/dgrijalva/jwt-go"

// Token is a struct containing Token data
type Token struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
