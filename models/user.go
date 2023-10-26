package models

import (
	"github.com/Bobby-P-dev/FinalProject2_kel7/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex" gorm:"not null"`                        //validation
	Email    string `json:"email" gorm:"uniqueIndex" gorm:"not null" validate:"required,email"` //validation
	Password string `json:"password" gorm:"not null" validate:"required,min=6"`                 //validation
	Age      int    `json:"age" gorm:"not null" validate:"required,numeric,min=18"`             //validation
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
