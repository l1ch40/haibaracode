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
	container.CID = dto.CID

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

func GetContainerByID(id uint) models.Container {
	container := models.Container{}
	models.DB.Find(&container, models.DB.Where("id = ?", id))
	return container
}

func CreateContainerProtocol(dto dto.ContainerProtocolDto) error {
	containerProtocol := models.ContainerProtocol{}
	containerProtocol.ContainerID = dto.CID
	containerProtocol.Protocol = dto.Protocol
	containerProtocol.HostIP = dto.HostIP
	containerProtocol.HostPort = dto.HostPort
	containerProtocol.ContainerPort = dto.ContainerPort

	err := models.DB.Create(&containerProtocol).Error
	return err
}

func DeleteContainerProtocol(dto dto.ContainerProtocolDto) error {
	containerProtocol := models.ContainerProtocol{}
	containerProtocol.ID = dto.ID
	if !HasContainerProtocol(dto) {
		return errors.New("项目中不存在该容器")
	}
	err := models.DB.Delete(&containerProtocol).Error
	return err
}

func CreateContainerAuth(dto dto.ContainerAuthDto) error {
	containerAuth := models.ContainerAuth{}
	containerAuth.ContainerID = dto.CID
	containerAuth.UserName = dto.Username
	containerAuth.Password = dto.Password

	err := models.DB.Create(&containerAuth).Error
	return err
}

func DeleteContainerAuth(dto dto.ContainerAuthDto) error {
	containerAuth := models.ContainerAuth{}
	containerAuth.ID = dto.ID
	if !HasContainerAuth(dto) {
		return errors.New("项目中不存在该容器")
	}
	err := models.DB.Delete(&containerAuth).Error
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

func HasContainerProtocol(dto dto.ContainerProtocolDto) bool {
	containerProtocol := models.ContainerProtocol{}
	result := models.DB.Where("id = ?", dto.ID).First(&containerProtocol)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func HasContainerAuth(dto dto.ContainerAuthDto) bool {
	containerAuth := models.ContainerAuth{}
	result := models.DB.Where("id = ?", dto.ID).First(&containerAuth)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}
