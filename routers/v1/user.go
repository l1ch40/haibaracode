package v1

import (
	"github.com/gin-gonic/gin"
	v1 "haibaracode/api/user/v1"
	"haibaracode/pkg/e"
)

func UserRouter(r *gin.Engine) {
	r.POST("/users", e.ErrorWrapper(v1.RegisterHandler))
	r.POST("/login", e.ErrorWrapper(v1.LoginHandle))
	r.DELETE("/sessions", e.ErrorWrapper(v1.LogoutHandle))
}
