package userservice

import responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"

type IUserService interface {
	ListAll() ([]responseDto.UserDto, error)
}
