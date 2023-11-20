package userservice

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	errortypes "CardozoCasariegoLuciano/StudyNotes/helpers/errorTypes"
	"sync"

	"github.com/devfeel/mapper"
)

var userS *userService
var once sync.Once

type userService struct {
	storage repository.IStorage
}

func NewUserService(storage repository.IStorage) IUserService {
	once.Do(func() {
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

func (userS *userService) FindByID(id uint) (*responseDto.UserDto, error) {
	user := userS.storage.GetUserByID(id)
	if user.CommonModelFields.ID == 0 {
		return nil, errortypes.UserNotFound
	}

	userDto := responseDto.UserDto{ID: user.CommonModelFields.ID}
	mapper.AutoMapper(&user, &userDto)

	return &userDto, nil
}

func (us *userService) EditUser(id uint, name string, image string) (*responseDto.UserDto, error) {

	userS.storage.EditUser(id, name, image)

	userByID, err := us.findByIdAndReturn(id)
	if err != nil {
		return nil, err
	}

	userDto := responseDto.UserDto{ID: userByID.CommonModelFields.ID}
	mapper.AutoMapper(userByID, &userDto)

	return &userDto, nil
}

func (userS *userService) findByIdAndReturn(id uint) (*models.User, error) {
	user := userS.storage.GetUserByID(id)
	if user.CommonModelFields.ID == 0 {
		return nil, errortypes.UserNotFound
	}

	return &user, nil
}
