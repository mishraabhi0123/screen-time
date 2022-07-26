package controllers

import (
	"github.com/dgrijalva/jwt-go"
	db "github.com/mishrabhi0123/screen-time/database"
	"github.com/mishrabhi0123/screen-time/models"
	"github.com/mishrabhi0123/screen-time/typedefs"
	"github.com/mishrabhi0123/screen-time/utils"
)

func RegisterUser(user models.User) typedefs.AppResponse {
	u := db.GetUser(user.Email)
	if len(u.Email) != 0 {
		return utils.CreateResponse(false, "This email is already registered with us", 400, nil)
	}
	user.Password, _ = utils.EncryptPassword(user.Password)
	user = db.CreateUser(user)
	return utils.CreateResponse(true, "Registration successful", 201, user)
}

func Login(currentUser models.User, deviceName string) typedefs.AppResponse {
	dbUser := db.GetUser(currentUser.Email)
	if len(dbUser.Email) == 0 {
		return utils.CreateResponse(false, "User not found", 404, nil)
	}

	match := utils.ComparePassword(currentUser.Password, dbUser.Password)
	if !match {
		return utils.CreateResponse(false, "Invalid Password", 400, nil)
	}

	device := db.GetDevice(dbUser.Id, deviceName)

	if device.Name == "" {
		device = db.RegisterDevice(deviceName, dbUser.Id)
	}

	tokenClaims := typedefs.TokenClaims{
		Email:          dbUser.Email,
		UserId:         dbUser.Id,
		DeviceName:     device.Name,
		StandardClaims: jwt.StandardClaims{},
	}
	signedToken := utils.CreateToken(tokenClaims)
	data := typedefs.Token{}
	data.Token = signedToken
	return utils.CreateResponse(true, "Logged in successfully", 200, data)
}
