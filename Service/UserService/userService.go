package userservice

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	errortypes "CardozoCasariegoLuciano/StudyNotes/helpers/errorTypes"
	"fmt"
	"sync"

	"github.com/devfeel/mapper"
)

var userS *userService
var once sync.Once

type userService struct {
	storage repository.IStorage
}

func NewUserService(storage repository.IStorage) *userService {
	once.Do(func() {
		fmt.Println("Pasa por aca userService")
		userS = &userService{storage: storage}
	})
	return userS
}

func (userS *userService) ListAll() ([]responseDto.UserDto, error) {
	allUsersDto := []responseDto.UserDto{}
	allUsersModels := []models.User{}
	err := userS.storage.ListAllUsers(&allUsersModels)
	if err != nil {
		return nil, errortypes.InternalError
	}

	for _, elem := range allUsersModels {
		userDto := responseDto.UserDto{ID: elem.CommonModelFields.ID}
		mapper.AutoMapper(&elem, &userDto)
		allUsersDto = append(allUsersDto, userDto)
	}

	return allUsersDto, nil
}
