package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"` //validation
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null"` //validation
	UsersID  int    `json:"user_id"`
	Users    Users
}
