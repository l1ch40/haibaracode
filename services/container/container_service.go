package container

import (
	"haibaracode/dto"
	container2 "haibaracode/repositories/container"
)

type ContainerService struct {
}

type ContainerProtocolService struct {
}

type ContainerAuthService struct {
}

func (c ContainerService) Create(cto dto.ContainerDto) error {
	return container2.CreateContainer(cto)
}

func (c ContainerService) Delete(cto dto.ContainerDto) error {
	return container2.DeleteContainer(cto)
}
