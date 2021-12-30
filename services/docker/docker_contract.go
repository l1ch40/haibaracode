package docker

import "haibaracode/dto"

type DockerContract interface {
	Create(dto dto.ContainerDto) (interface{}, error)
	Delete(dto dto.ContainerDto) error
}
