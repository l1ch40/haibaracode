package api

import (
	"github.com/gin-gonic/gin"
	"haibaracode/pkg/cookie"
	"haibaracode/pkg/e"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := cookie.ReadCookie(c)
		if err != nil {
			apiError := e.ApiError{
				Status:  http.StatusUnauthorized,
				Code:    401,
				Message: "认证失败",
			}
			c.JSON(apiError.Status, apiError)
			return
		}
		c.Next()
	}
}
