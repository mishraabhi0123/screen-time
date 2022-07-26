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

func GetDevice(userId int, deviceName string) models.Device {
	devices := Devices[userId]
	foundDevice := models.Device{UserId: userId}

	for _, device := range devices {
		if device.Name == deviceName {
			foundDevice = device
			break
		}
	}

	return foundDevice
}

func RegisterDevice(deviceName string, userId int) models.Device {
	devices := Devices[userId]
	device := models.Device{}

	device.UserId = userId
	device.Id = len(devices) + 1
	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()

	devices = append(devices, device)
	Devices[device.UserId] = devices
	return device
}
