package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"haibaracode/dto"
	"haibaracode/pkg/cookie"
	"haibaracode/pkg/e"
	"haibaracode/requests/v1"
	"haibaracode/services/user"
)

func RegisterHandler(c *gin.Context) (interface{}, error) {
	request := v1.RegisterRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}

	userDto := dto.UserDto{
		Username: request.Username,
		Password: request.Password,
		UID:      uuid.New().String(),
	}

	service := user.UserService{}
	err = service.Register(userDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}

	return "注册成功", nil
}

func LoginHandle(c *gin.Context) (interface{}, error) {
	request := v1.LoginRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}

	userDto := dto.UserDto{
		Username: request.Username,
		Password: request.Password,
	}

	service := user.UserService{}
	err = service.Login(userDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}

	userID, err := service.GetUserIDByUsername(userDto.Username)
	err = cookie.CreateCookie(c, userID)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40006,
			Message: err.Error(),
		}
	}
	return "登录成功", nil
}

func LogoutHandle(c *gin.Context) (interface{}, error) {
	userID, err := cookie.ReadCookie(c)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40007,
			Message: err.Error(),
		}
	}
	service := user.UserService{}
	if service.HasUserID(userID) {
		cookie.DeleteCookie(c)
		return "登出成功。", nil
	}
	return nil, errors.New("登出失败。")
}
