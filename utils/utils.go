package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mishrabhi0123/screen-time/typedefs"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte = []byte("ashdjahsdj")

func CreateResponse(success bool, message string, statusCode int, data any) typedefs.AppResponse {
	var res typedefs.AppResponse
	res.Success = success
	res.Message = message
	res.StatusCode = statusCode
	res.Data = data
	return res
}

func Print(object any) {
	fmt.Printf("%+v\n", object)
}

func EncryptPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordBytes), err
}

func ComparePassword(claimedPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(claimedPassword))
	return err == nil
}

func SendResponse(w http.ResponseWriter, controllerResponse typedefs.AppResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(controllerResponse.StatusCode)
	jsonBytes, _ := json.Marshal(controllerResponse)
	w.Write(jsonBytes)
}

func CreateToken(claims typedefs.TokenClaims) string {
	// unsigned token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with a secret key
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	return signedToken
}

func ParseToken(token string) (typedefs.TokenClaims, error) {
	tokenClaims := typedefs.TokenClaims{}

	tkn, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return tokenClaims, err
	}
	if !tkn.Valid {
		return tokenClaims, errors.New("invalid token")
	}

	return tokenClaims, nil
}
