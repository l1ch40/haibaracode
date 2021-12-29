package models

import "github.com/jinzhu/gorm"

type Container struct {
	gorm.Model
	UID    uint `gorm:"not null"`
	Status uint
	Image  string
}

type ContainerAuth struct {
	gorm.Model
	ContainerID uint `gorm:"not null"`
	UserName    string
	Password    string
}

type ContainerProtocol struct {
	gorm.Model
	ContainerID   uint `gorm:"not null"`
	Protocol      string
	HostIP        string
	HostPort      uint
	ContainerPort uint
}
