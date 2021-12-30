package v1

import (
	"github.com/gin-gonic/gin"
	"haibaracode/dto"
	"haibaracode/pkg/cookie"
	"haibaracode/pkg/e"
	v1 "haibaracode/requests/v1"
	"haibaracode/services/project"
)

func CreateHandle(c *gin.Context) (interface{}, error) {
	request := v1.CreateProjectRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}
	userID, err := cookie.ReadCookie(c)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40007,
			Message: err.Error(),
		}
	}

	projectDto := dto.ProjectDto{
		Name: request.Name,
		UID:  userID,
	}

	service := project.ProjectService{}
	err = service.Create(projectDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}
	return "项目创建成功", nil
}

func DeleteHandle(c *gin.Context) (interface{}, error) {
	request := v1.DeleteProjectRequest{}

	err := c.ShouldBindUri(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}

	projectDto := dto.ProjectDto{
		ID: request.ID,
	}
	service := project.ProjectService{}
	err = service.Delete(projectDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}
	return "项目删除成功", nil
}
