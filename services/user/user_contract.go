package user

import "haibaracode/dto"

type UserContract interface {
	Register(dto dto.UserDto) error
	Login(dto dto.UserDto) error
}
