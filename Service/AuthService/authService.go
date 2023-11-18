package authservice

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	errortypes "CardozoCasariegoLuciano/StudyNotes/helpers/errorTypes"
	"CardozoCasariegoLuciano/StudyNotes/helpers/roles"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"
	"fmt"
	"sync"

	"github.com/devfeel/mapper"
)

var authS *authService
var once sync.Once

type authService struct {
	storage    repository.IStorage
	encripting utils.Ibcrypt
}

func NewAuthService(storage repository.IStorage, cripto utils.Ibcrypt) *authService {
	once.Do(func() {
		fmt.Println("Pasa por aca authService dentro del once")
		authS = &authService{
			storage:    storage,
			encripting: cripto,
		}
	})
	return authS
}

func (auth *authService) RegisterUser(user requestDto.RegisterUserDto) (*responseDto.UserDto, error) {
	//Validate email
	userEmail := auth.storage.FindUserByEmail(user.Email)
	if userEmail.ID != 0 {
		return nil, errortypes.MailAlreadyTaken
	}

	//Hashing ths password
	hashedPass, err := auth.encripting.HashPassword(user.Password)
	if err != nil {
		return nil, errortypes.InternalError
	}

	userM := models.User{Role: roles.USER}
	mapper.AutoMapper(&user, &userM)

	//Save user
	userM.Password = string(hashedPass)
	err = auth.storage.SaveUser(&userM)
	if err != nil {
		return nil, errortypes.InternalError
	}

	userDto := responseDto.UserDto{ID: userM.CommonModelFields.ID}
	mapper.AutoMapper(&userM, &userDto)

	return &userDto, nil
}

func (auth *authService) LoginUser(user requestDto.LoginUserDto) (*responseDto.UserDto, error) {
	userLoged := auth.storage.FindUserByEmail(user.Email)
	if userLoged.ID == 0 {
		return nil, errortypes.WrongPassOrEmail
	}

	//Compare the passwords
	err := auth.encripting.Compare(userLoged.Password, user.Password)
	if err != nil {
		return nil, errortypes.WrongPassOrEmail
	}

	userDto := responseDto.UserDto{ID: userLoged.CommonModelFields.ID}
	mapper.AutoMapper(&userLoged, &userDto)
	return &userDto, nil
}
