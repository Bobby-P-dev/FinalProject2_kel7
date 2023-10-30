package models

import (
	"github.com/Bobby-P-dev/FinalProject2_kel7/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"uniqueIndex" gorm:"not null"`                     //validation
	Email    string `json:"email" form:"email" gorm:"uniqueIndex" gorm:"not null" validate:"required,email"` //validation
	Password string `json:"password" gorm:"not null" validate:"required,min=6"`                              //validation
	Age      int    `json:"age" gorm:"not null" validate:"required,numeric,min=18"`                          //validation
	PhotoUrl string `json:"photo_url" form:"photo_url"`
	Photo    photoResponeU
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

type UsersRespon struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username" `
	Email    string `json:"email" form:"email"`
}

func (UsersRespon) TableName() string {
	return "users"
}

type UsersResponSocial struct {
	ID       int    `json:"id" form:"users_id"`
	Username string `json:"username" form:"username"`
	PhotoUrl string `json:"profile_image_url" form:"photo_url"`
}

func (UsersResponSocial) TableName() string {
	return "users"
}
