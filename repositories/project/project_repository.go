package project

import (
	"haibaracode/dto"
	"haibaracode/models"
)

func CreateProject(dto dto.ProjectDto) error {
	project := models.Project{}
	project.UID = dto.UID
	project.Name = dto.Name

	err := models.DB.Create(&project).Error
	return err
}

func DeleteProject(dto dto.ProjectDto) error {
	project := models.Project{}
	project.ID = dto.ID
	project.UID = dto.UID
	project.Name = dto.Name

	err := models.DB.Delete(&project).Error
	return err
}
