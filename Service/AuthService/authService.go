package authservice

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	"fmt"
	"sync"

	"github.com/devfeel/mapper"
)

var authS *authService
var once sync.Once

type authService struct {
	storage repository.IStorage
}

func NewAuthService() *authService {
	once.Do(func() {
		fmt.Println("Pasa por aca authService dentro del once")
		authS = &authService{storage: repository.NewMemory()}
	})
	return authS
}

func (auth *authService) RegisterUser(user requestDto.RegisterUserDto) responseDto.ResponseDto {
	userM := models.User{Role: "User"}
	mapper.AutoMapper(&user, &userM)

	//TODO hacer todo el tema de los JWT
	//TODO Tirar un error si ya existe o si surgio un error Seguir por aca
	savedUser := auth.storage.Save(userM)

	userDto := responseDto.UserDto{}
	mapper.AutoMapper(&savedUser, &userDto)

	resp := responseDto.NewResponse("OK", "Usuario creado", userDto)
	return resp
}
