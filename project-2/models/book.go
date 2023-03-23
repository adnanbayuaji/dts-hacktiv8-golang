package models

type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null; type:varchar(191)"`
	Author      string `gorm:"not null; type:varchar(191)"`
	Description string `gorm:"not null; type:varchar(191)"`
}
