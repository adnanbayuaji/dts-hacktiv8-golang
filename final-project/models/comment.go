package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an comment
type Comment struct {
	GormModel
	Message string `json:"message" form:"message" valid:"required~Message is required"`
	UserID  uint
	User    *User
	PhotoID uint
	Photo   *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
