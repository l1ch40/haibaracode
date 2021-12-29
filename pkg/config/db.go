package config

import (
	"gopkg.in/ini.v1"
	"haibaracode/pkg/helper"
)

var MySQL *mysql

type mysql struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`

	source *ini.File
}

func (s *mysql) Load(path string) *mysql {
	var err error

	exists, err := helper.PathExists(path)
	if !exists {
		return s
	}
	s.source, err = ini.Load(path)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *mysql) Init() *mysql {
	if s.source == nil {
		return s
	}

	err := s.source.Section("mysql").MapTo(s)
	if err != nil {
		panic(err)
	}
	return s
}
