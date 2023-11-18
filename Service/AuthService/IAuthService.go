package authservice

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
)

type IAuthService interface {
	RegisterUser(user requestDto.RegisterUserDto) (*responseDto.UserDto, error)
	LoginUser(user requestDto.LoginUserDto) (*responseDto.UserDto, error)
}
