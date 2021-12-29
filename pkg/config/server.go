package config

import (
	"gopkg.in/ini.v1"
	"haibaracode/pkg/helper"
)

var Server *server

type server struct {
	Address string
	Port    int
	source  *ini.File
}

func (s *server) Load(path string) *server {
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

func (s *server) Init() *server {
	if s.source == nil {
		return s
	}
	s.Address = s.source.Section("server").Key("address").MustString("0.0.0.0")
	s.Port = s.source.Section("server").Key("port").MustInt(8080)
	return s
}
