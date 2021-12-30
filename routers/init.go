package routers

import (
	"github.com/gin-gonic/gin"
	v1 "haibaracode/routers/v1"
)

func Init(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	v1.UserRouter(r)
	v1.ProjectRouter(r)
}
