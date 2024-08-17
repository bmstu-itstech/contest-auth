package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

type UserInfoJWT struct {
	Email string `json:"email"`
}
