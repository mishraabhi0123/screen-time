package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mishrabhi0123/screen-time/typedefs"
	"golang.org/x/crypto/bcrypt"
)

func CreateControllerResponse(success bool, message string, statusCode int, data any) typedefs.ControllerResponse {
	var res typedefs.ControllerResponse
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

func SendResponse(w http.ResponseWriter, controllerResponse typedefs.ControllerResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(controllerResponse.StatusCode)
	jsonBytes, _ := json.Marshal(controllerResponse)
	w.Write(jsonBytes)
}
