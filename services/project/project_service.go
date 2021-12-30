package project

import (
	"haibaracode/dto"
	project2 "haibaracode/repositories/project"
)

type ProjectService struct {
}

func (project ProjectService) Create(projectDto dto.ProjectDto) error {
	return project2.CreateProject(projectDto)
}

func (project ProjectService) Delete(projectDto dto.ProjectDto) error {
	return project2.DeleteProject(projectDto)
}
