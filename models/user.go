package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UId      string `gorm:"unique;not null"`
	Username string
	Password string `gorm:"type:varchar(100);not null"`
}
