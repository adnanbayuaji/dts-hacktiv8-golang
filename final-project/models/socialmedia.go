package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Your username is required"`
	SocialMediaUrl string `gorm:"not null" json:"socialmediaurl" form:"socialmediaurl" valid:"required~Your username is required"`
	UserID         uint
	User           *User
}

func (s *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
