package v1

import (
	"github.com/gin-gonic/gin"
	"haibaracode/api"
	containerV1 "haibaracode/api/container/v1"
	projectV1 "haibaracode/api/project/v1"
	"haibaracode/pkg/e"
)

func ProjectRouter(r *gin.Engine) {
	authorizedProject := r.Group("/")
	authorizedProject.Use(api.AuthRequired())
	{
		authorizedProject.POST("/projects", e.ErrorWrapper(projectV1.CreateHandle))
		authorizedProject.DELETE("/projects/:id", e.ErrorWrapper(projectV1.DeleteHandle))
		authorizedProject.POST("/projects/:id/containers", e.ErrorWrapper(containerV1.CreateHandle))
		authorizedProject.DELETE("/projects/:id/containers/:cid", e.ErrorWrapper(containerV1.DeleteHandle))
	}

}
