package database

import (
	"time"

	"github.com/mishrabhi0123/screen-time/models"
)

func GetUser(email string) models.User {
	user := Users[email]
	return user
}

func CreateUser(user models.User) models.User {
	user.Id = len(Users) + 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	Users[user.Email] = user
	return user
}

func GetDevice(deviceId int) models.Device {
	device := Devices[deviceId]
	return device
}
