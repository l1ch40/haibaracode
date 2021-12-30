package project

import "haibaracode/dto"

type ProjectContract interface {
	Create(dto dto.ProjectDto) error
	Delete(dto dto.ProjectDto) error
}
