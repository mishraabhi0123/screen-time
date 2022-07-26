package database

import "github.com/mishrabhi0123/screen-time/models"

var Users = map[string]models.User{}
var Devices = map[int][]models.Device{}
