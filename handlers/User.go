package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mishrabhi0123/screen-time/controllers"
	"github.com/mishrabhi0123/screen-time/models"
	"github.com/mishrabhi0123/screen-time/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(bodyBytes, &user)
	controllerResponse := controllers.RegisterUser(user)
	utils.SendResponse(w, controllerResponse)
}

func Login(w http.ResponseWriter, r *http.Request) {
	deviceName := r.Header.Get("x-device-name")
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	var currentUser models.User
	json.Unmarshal(bodyBytes, &currentUser)
	controllerResponse := controllers.Login(currentUser, deviceName)
	utils.SendResponse(w, controllerResponse)
}
