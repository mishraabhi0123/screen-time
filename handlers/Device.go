package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mishrabhi0123/screen-time/controllers"
	"github.com/mishrabhi0123/screen-time/utils"
)

func GetDevice(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	paths := strings.Split(url, "/")
	id := paths[len(paths)-1]
	deviceId, err := strconv.Atoi(id)
	if err != nil {
		res := utils.CreateControllerResponse(false, "Invalid device id", 400, nil)
		utils.SendResponse(w, res)
		return
	}

	controllerResponse := controllers.GetDevice(deviceId)
	utils.SendResponse(w, controllerResponse)
}
