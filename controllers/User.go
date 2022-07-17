package controllers

import (
	db "github.com/mishrabhi0123/screen-time/database"
	"github.com/mishrabhi0123/screen-time/models"
	"github.com/mishrabhi0123/screen-time/utils"
)

func RegisterUser(user models.User) utils.ControllerResponse {
	u := db.GetUser(user.Email)
	if len(u.Email) != 0 {
		return utils.CreateControllerResponse(false, "This email is already registered with us", 400, nil)
	}
	user.Password, _ = utils.EncryptPassword(user.Password)
	user = db.CreateUser(user)
	return utils.CreateControllerResponse(true, "Registration successful", 201, user)
}

func InitiateLogin() utils.ControllerResponse {
	return utils.CreateControllerResponse(true, "Registration successful", 201, nil)
}

func Login() utils.ControllerResponse {
	return utils.CreateControllerResponse(true, "Registration successful", 201, nil)
}
