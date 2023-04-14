package models

import (
	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string    `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your email is required,email~Invalid email format"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age string    `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(8|100)~Age has to have a minimum at 8 years old"`
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
