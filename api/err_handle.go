package api

import (
	"github.com/gin-gonic/gin"
	"haibaracode/pkg/e"
	"net/http"
)

func ErrorHandler(c *gin.Context) (interface{}, error) {
	query := c.Query("q")
	if query == "" {
		return nil, e.ApiError{
			Status:  http.StatusOK,
			Code:    404,
			Message: "q 的参数不能为空",
		}
	}

	if query == "test" {
		return nil, e.ApiError{
			Status:  http.StatusOK,
			Code:    404,
			Message: "q 的参数不能为 test",
		}
	}
	return query, nil
}
