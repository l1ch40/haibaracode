package container

import (
	"errors"
	"haibaracode/dto"
	"haibaracode/models"
)

func CreateContainer(dto dto.ContainerDto) error {
	container := models.Container{}
	container.PID = dto.PID
	container.Image = dto.Image
	container.Status = dto.Status

	err := models.DB.Create(&container).Error
	return err
}

func DeleteContainer(dto dto.ContainerDto) error {
	container := models.Container{}
	container.ID = dto.ID
	if !ContainerInProject(dto) {
		return errors.New("项目中不存在该容器")
	}
	err := models.DB.Delete(&container).Error
	return err
}

func ContainerInProject(dto dto.ContainerDto) bool {
	container := models.Container{}
	result := models.DB.Where("id = ? AND p_id = ?", dto.ID, dto.PID).First(&container)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}
