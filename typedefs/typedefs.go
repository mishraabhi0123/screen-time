package typedefs

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type AppResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}

type TokenClaims struct {
	Email      string `json:"email"`
	UserId     int    `json:"userId"`
	DeviceName string `json:"deviceId"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeviceInfo struct {
	Name            string `json:"deviceName"`
	OperatingSystem string `json:"operatingSystem"`
}
