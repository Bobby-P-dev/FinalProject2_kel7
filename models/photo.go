package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title" form:"title" gorm:"not null"` //validation
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" gorm:"not null"` //validation
	UsersID  int    `json:"users_id" form:"users_id"`
	Users    UsersRespon
}

type photoRespone struct {
	ID       int    `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
	UsersID  int    `json:"users_id" form:"users_id"`
}

func (photoRespone) TableName() string {
	return "photos"
}

type photoResponeU struct {
	ID       int    `json:"id" form:"id"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
}

func (photoResponeU) TableName() string {
	return "photos"
}
