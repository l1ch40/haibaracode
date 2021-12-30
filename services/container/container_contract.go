package container

import "haibaracode/dto"

type ContainerContract interface {
	Create(dto dto.ContainerDto) error
	Delete(dto dto.ContainerDto) error
}

type ContainerProtocolContract interface {
	Create(dto dto.ContainerProtocolDto) error
	Delete(dto dto.ContainerProtocolDto) error
}

type ContainerAuthContract interface {
	Create(dto dto.ContainerAuthDto) error
	Delete(dto dto.ContainerAuthDto) error
}
