package user

import (
	"errors"
	"haibaracode/dto"
	"haibaracode/pkg/hash"
	user2 "haibaracode/repositories/user"
)

type UserService struct {
}

func (user UserService) Register(userDto dto.UserDto) error {
	bytes, err := hash.NewHash().Make([]byte(userDto.Password))
	if err != nil {
		return errors.New(err.Error())
	}
	userDto.Password = string(bytes)
	return user2.CreateUser(userDto)
}

func (user UserService) Login(userDto dto.UserDto) error {
	model := user2.GetUserByUsername(userDto.Username)
	if model.ID == 0 {
		return errors.New("账号不存在")
	}
	err := hash.NewHash().Check([]byte(model.Password), []byte(userDto.Password))
	if err != nil {
		return errors.New("密码错误")
	}

	return nil
}

func (user UserService) GetUserIDByUsername(userDto dto.UserDto) (string, error) {
	model := user2.GetUserByUsername(userDto.Username)
	if model.ID == 0 {
		return "", errors.New("账号不存在")
	}
	return model.UID, nil

}
