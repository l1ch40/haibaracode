package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"haibaracode/pkg/config"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database)

	DB, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	migrate()
}

func migrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Project{})
	DB.AutoMigrate(&Container{})
	DB.AutoMigrate(&ContainerAuth{})
	DB.AutoMigrate(&ContainerProtocol{})
}
