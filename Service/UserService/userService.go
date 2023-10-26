package userservice

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	"fmt"
	"sync"
)

var userS *userService
var once sync.Once

type userService struct {
	storage repository.IMemory
}

func NewAuthService() *userService {
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
		temp := responseDto.UserDto{
			Name:  elem.Name,
			Id:    elem.Id,
			Email: elem.Email,
			Image: elem.Image,
			Role:  elem.Role,
		}
		allUsersDto = append(allUsersDto, temp)
	}

	resp := responseDto.NewResponse("OK", "Lista de usuarios", allUsersDto)
	//TODO ver si tendria que crear un "ObjectMapper"
	//TODO hacer esto con go routines y ver que tanto cambia el tiempo
	return resp
}
