package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func (user *User) ToJson() (string, error) {
	u, err := json.Marshal(user)
	return string(u), err
}
