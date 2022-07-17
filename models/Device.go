package models

import "time"

type Device struct {
	Id              int       `json:"id"`
	UserId          int       `json:"userId"`
	Name            string    `json:"name"`
	OperatingSystem string    `json:"operatingSystem"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	DeletedAt       time.Time `json:"deletedAt"`
}
