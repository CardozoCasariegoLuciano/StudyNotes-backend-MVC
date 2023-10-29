package userservice

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	"fmt"
	"sync"

	"github.com/devfeel/mapper"
)

var userS *userService
var once sync.Once

type userService struct {
	storage repository.IStorage
}

func NewUserService() *userService {
	once.Do(func() {
		fmt.Println("Pasa por aca userService")
		userS = &userService{storage: repository.NewMemory()}
	})
	return userS
}

func (userS *userService) ListAll() responseDto.ResponseDto {
	allusers := userS.storage.ListAll()
	allUsersDto := []responseDto.UserDto{}

	for _, elem := range allusers {
		userDto := responseDto.UserDto{}
		mapper.AutoMapper(&elem, &userDto)
		allUsersDto = append(allUsersDto, userDto)
	}

	resp := responseDto.NewResponse("OK", "Lista de usuarios", allUsersDto)
	return resp
}
