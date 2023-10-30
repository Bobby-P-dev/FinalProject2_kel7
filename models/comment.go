package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"not null"` //validation
	PhotoID int    `json:"photo_id" form:"photo_id"`
	UsersID int    `json:"users_id" form:"users_id"`
	Users   UsersRespon
	Photo   photoRespone
}
