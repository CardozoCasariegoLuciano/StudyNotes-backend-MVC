package registerdto

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	"CardozoCasariegoLuciano/StudyNotes/helpers/roles"
)

// RegisterDto
var RegisterEmailTaken requestDto.RegisterUserDto = requestDto.RegisterUserDto{
	Name:         "Register",
	Email:        "test@mailTaken.com",
	Password:     "taken",
	Confirmation: "taken",
}

var RegisterInternalError requestDto.RegisterUserDto = requestDto.RegisterUserDto{
	Name:         "Register",
	Email:        "Internal@error.com",
	Password:     "internal",
	Confirmation: "internal",
}

var RegisterSucces requestDto.RegisterUserDto = requestDto.RegisterUserDto{
	Name:         "Register",
	Email:        "Succes@succes.com",
	Password:     "succes",
	Confirmation: "succes",
}

// UserDto
var UserRegistered responseDto.UserDto = responseDto.UserDto{
	ID:    2,
	Name:  RegisterSucces.Name,
	Email: RegisterSucces.Email,
	Role:  roles.USER,
}
