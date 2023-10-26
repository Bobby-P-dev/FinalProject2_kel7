package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UsersID int `json:"user_id"`
	Users   Users
	PhotoID int `json:"photo_id"`
	Photo   Photo
	Message string `json:"message" gorm:"not null"` //validation
}
