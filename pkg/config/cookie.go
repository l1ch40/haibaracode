package config

import (
	"gopkg.in/ini.v1"
	"haibaracode/pkg/helper"
)

var Cookie *cookie

type cookie struct {
	HashKey  string
	BlockKey string
	source   *ini.File
}

func (c *cookie) Load(path string) *cookie {
	var err error

	exists, err := helper.PathExists(path)
	if !exists {
		return c
	}
	c.source, err = ini.Load(path)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *cookie) Init() *cookie {
	if c.source == nil {
		return c
	}
	c.HashKey = c.source.Section("cookie").Key("hash_key").MustString("123")
	c.BlockKey = c.source.Section("cookie").Key("block_key").MustString("321")
	return c
}
