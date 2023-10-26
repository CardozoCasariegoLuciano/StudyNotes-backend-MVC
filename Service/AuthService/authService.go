package authservice

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	"fmt"
	"sync"
)

var authS *authService
var once sync.Once

type authService struct {
	storage repository.IMemory
}

func NewAuthService() *authService {
	once.Do(func() {
		fmt.Println("Pasa por aca authService")
		authS = &authService{storage: repository.NewMemory()}
	})
	return authS
}

func (auth *authService) RegisterUser(user requestDto.RegisterUserDto) responseDto.ResponseDto {
	userM := models.User{
		Name:     user.Name,
		Role:     "USER",
		Email:    user.Email,
		Image:    "",
		Password: user.Password,
	}

	//TODO Tirar un error si ya existe o si surgio un error
	savedUser := auth.storage.Save(userM)

	userDto := responseDto.UserDto{
		Id:    savedUser.Id,
		Name:  savedUser.Name,
		Email: savedUser.Email,
		Image: savedUser.Image,
		Role:  savedUser.Role,
	}

	resp := responseDto.NewResponse("OK", "Usuario creado", userDto)
	return resp
}
