package routers

import (
	"github.com/gin-gonic/gin"
	v1 "haibaracode/routers/v1"
)

func Init(r *gin.Engine) {
	v1.UserRouter(r)
}
