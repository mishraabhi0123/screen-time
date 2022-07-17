package controllers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	db "github.com/mishrabhi0123/screen-time/database"
	"github.com/mishrabhi0123/screen-time/models"
	"github.com/mishrabhi0123/screen-time/typedefs"
	"github.com/mishrabhi0123/screen-time/utils"
)

func RegisterUser(user models.User) typedefs.ControllerResponse {
	u := db.GetUser(user.Email)
	if len(u.Email) != 0 {
		return utils.CreateControllerResponse(false, "This email is already registered with us", 400, nil)
	}
	user.Password, _ = utils.EncryptPassword(user.Password)
	user = db.CreateUser(user)
	return utils.CreateControllerResponse(true, "Registration successful", 201, user)
}

func InitiateLogin() typedefs.ControllerResponse {
	return utils.CreateControllerResponse(true, "Registration successful", 201, nil)
}

func Login(currentUser models.User) typedefs.ControllerResponse {
	dbUser := db.GetUser(currentUser.Email)
	if len(dbUser.Email) == 0 {
		return utils.CreateControllerResponse(false, "User not found", 404, nil)
	}
	match := utils.ComparePassword(currentUser.Password, dbUser.Password)
	if !match {
		return utils.CreateControllerResponse(false, "Invalid Password", 400, nil)
	}

	tokenClaims := typedefs.TokenClaims{
		Email:          dbUser.Email,
		StandardClaims: jwt.StandardClaims{},
	}
	// unsigned token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	// sign the token with a secret key
	signedToken, err := token.SignedString([]byte("ashdjahsdj"))
	if err != nil {
		fmt.Println(err.Error())
	}
	data := typedefs.Token{}
	data.Token = signedToken
	return utils.CreateControllerResponse(true, "Logged in successfully", 200, data)
}
