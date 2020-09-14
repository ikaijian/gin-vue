package dto

import "github.com/kaijian/gin-vue/model"

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto  {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}
