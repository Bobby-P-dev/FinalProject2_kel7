package models

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`             //validation
	SocialMediaUrl string `json:"social_media_url" gorm:"not null"` //validation
	UsersID        int    `json:"users_id" form:"user_id"`
	Users          UsersResponSocial
}
