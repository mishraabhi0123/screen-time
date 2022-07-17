package typedefs

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type ControllerResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}

type TokenClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token"`
}
