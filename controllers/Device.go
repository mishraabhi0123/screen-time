package controllers

import (
	"github.com/mishrabhi0123/screen-time/database"
	"github.com/mishrabhi0123/screen-time/typedefs"
	"github.com/mishrabhi0123/screen-time/utils"
)

func GetDevice(deviceId int) typedefs.ControllerResponse {
	device := database.GetDevice(deviceId)
	if device.Id != deviceId {
		return utils.CreateControllerResponse(false, "Could not find any device with this id", 404, nil)
	}
	return utils.CreateControllerResponse(true, "Device found", 200, device)
}
