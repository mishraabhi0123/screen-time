package controllers

import (
	"github.com/mishrabhi0123/screen-time/database"
	"github.com/mishrabhi0123/screen-time/typedefs"
	"github.com/mishrabhi0123/screen-time/utils"
)

func GetDevice(userId int, deviceName string) typedefs.AppResponse {
	device := database.GetDevice(userId, deviceName)
	if device.Id == 0 {
		return utils.CreateResponse(false, "Could not find any device with this id", 404, nil)
	}
	return utils.CreateResponse(true, "Device found", 200, device)
}
