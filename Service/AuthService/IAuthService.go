package authservice

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
)

type IAuthService interface {
	RegisterUser(user requestDto.RegisterUserDto) (responseDto.ResponseDto, int)
}
