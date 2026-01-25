package models

import jwt "github.com/dgrijalva/jwt-go"

// JWTClaims describes Claims in token.
type JWTClaims struct {
	ID       string `json:"id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

//NewJWTClaims initialize tokens
func (claims JWTClaims) NewJWTClaims(userType, id string) JWTClaims {

	return JWTClaims{
		ID:       id,
		UserType: userType,
	}
}
