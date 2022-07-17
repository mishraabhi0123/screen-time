package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type ControllerResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}

func CreateControllerResponse(success bool, message string, statusCode int, data any) ControllerResponse {
	res := ControllerResponse{success, message, statusCode, data}
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

func SendResponse(w http.ResponseWriter, controllerResponse ControllerResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(controllerResponse.StatusCode)
	jsonBytes, _ := json.Marshal(controllerResponse)
	w.Write(jsonBytes)
}
