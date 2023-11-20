package userservice

import responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"

type IUserService interface {
	ListAll() ([]responseDto.UserDto, error)
	FindByID(id uint) (*responseDto.UserDto, error)
}
