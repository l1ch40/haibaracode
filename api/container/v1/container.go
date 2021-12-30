package v1

import (
	"github.com/gin-gonic/gin"
	"haibaracode/dto"
	"haibaracode/pkg/e"
	v1 "haibaracode/requests/v1"
	"haibaracode/services/container"
)

func CreateHandle(c *gin.Context) (interface{}, error) {
	pidRequest := v1.CreateContainerRequestUri{}
	errPID := c.ShouldBindUri(&pidRequest)
	imageRequest := v1.CreateContainerRequestJson{}
	errImage := c.ShouldBind(&imageRequest)

	if errPID != nil || errImage != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}

	containerDto := dto.ContainerDto{}
	containerDto.PID = pidRequest.PID
	containerDto.Image = imageRequest.Image

	service := container.ContainerService{}
	err := service.Create(containerDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}
	return "容器创建成功", nil
}

func DeleteHandle(c *gin.Context) (interface{}, error) {
	request := v1.DeleteContainerRequest{}

	err := c.ShouldBindUri(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}

	containerDto := dto.ContainerDto{}
	containerDto.ID = request.ID
	containerDto.PID = request.PID

	service := container.ContainerService{}
	err = service.Delete(containerDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}
	return "容器删除成功", nil
}
