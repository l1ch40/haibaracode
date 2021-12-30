package models

import "github.com/jinzhu/gorm"

type Container struct {
	gorm.Model
	PID    uint `gorm:"not null"`
	CID    string
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
	HostPort      string
	ContainerPort string
}
