package container

import (
	"haibaracode/dto"
	container2 "haibaracode/repositories/container"
	"haibaracode/services/docker"
)

const (
	STARTED = iota
)

type ContainerService struct {
}

func (c ContainerService) Create(cto dto.ContainerDto) error {
	dockerService := docker.DockerService{}
	containerID, err := dockerService.Create(cto)
	if err != nil {
		return err
	}
	cto.CID = containerID.(string)
	cto.Status = STARTED
	return container2.CreateContainer(cto)
}

func (c ContainerService) Delete(cto dto.ContainerDto) error {
	container := container2.GetContainerByID(cto.ID)
	cto.CID = container.CID
	dockerService := docker.DockerService{}
	if err := dockerService.Delete(cto); err != nil {
		return err
	}
	return container2.DeleteContainer(cto)
}

type ContainerProtocolService struct {
}

func (c ContainerProtocolService) Create(cto dto.ContainerProtocolDto) error {
	return container2.CreateContainerProtocol(cto)
}

func (c ContainerProtocolService) Delete(cto dto.ContainerProtocolDto) error {
	return container2.DeleteContainerProtocol(cto)
}

type ContainerAuthService struct {
}

func (c ContainerAuthService) Create(cto dto.ContainerAuthDto) error {
	return container2.CreateContainerAuth(cto)
}

func (c ContainerAuthService) Delete(cto dto.ContainerAuthDto) error {
	return container2.DeleteContainerAuth(cto)
}
