package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "haibaracode/models"
	"haibaracode/pkg/config"
	"haibaracode/routers"
)

func main() {
	r := gin.New()
	routers.Init(r)
	r.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port))
}
