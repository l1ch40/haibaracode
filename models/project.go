package models

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	UID  string `gorm:"not null"`
	Name string `gorm:"not null"`
}
