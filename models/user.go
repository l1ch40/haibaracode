package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UID      string `gorm:"unique;not null"`
	Username string
	Password string `gorm:"type:varchar(100);not null"`
}
