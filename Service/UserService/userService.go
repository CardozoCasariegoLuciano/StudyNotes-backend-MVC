package userservice

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	mysql "CardozoCasariegoLuciano/StudyNotes/Repository/MySql"
	"fmt"
	"net/http"
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
		userS = &userService{storage: mysql.NewDataBase()}
	})
	return userS
}

func (userS *userService) ListAll() (responseDto.ResponseDto, int) {
	allUsersDto := []responseDto.UserDto{}
	allUsersModels := []models.User{}
	userS.storage.ListAllUsers(&allUsersModels)

	for _, elem := range allUsersModels {
		userDto := responseDto.UserDto{ID: int(elem.CommonModelFields.ID)}
		mapper.AutoMapper(&elem, &userDto)
		allUsersDto = append(allUsersDto, userDto)
	}

	resp := responseDto.NewResponse("OK", "Lista de usuarios", allUsersDto)
	return resp, http.StatusOK
}
