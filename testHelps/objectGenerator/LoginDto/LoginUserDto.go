package logindto

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	"CardozoCasariegoLuciano/StudyNotes/helpers/roles"
)

// LoginUserDto
var LoginEmailDontExist requestDto.LoginUserDto = requestDto.LoginUserDto{
	Password: "EmailDontexist",
	Email:    "EmailDontexist@Test.com",
}

var LoginInternalError requestDto.LoginUserDto = requestDto.LoginUserDto{
	Password: "InternalError",
	Email:    "Internal@error.com",
}

var LoginSuccesfuly requestDto.LoginUserDto = requestDto.LoginUserDto{
	Password: "succes",
	Email:    "succes@login.com",
}

var LoginHasEmailButPassDontMatch requestDto.LoginUserDto = requestDto.LoginUserDto{
	Password: "dontMatch",
	Email:    "valid@email.com",
}

// UserDto
var UserLoged responseDto.UserDto = responseDto.UserDto{
	ID:    22,
	Email: LoginSuccesfuly.Email,
	Role:  roles.ADMIN,
	Image: "loged",
	Name:  "success",
}

// Model.User
var UserInStorageToLogin models.User = models.User{
	Email: LoginSuccesfuly.Email,
	Role:  roles.ADMIN,
	Image: "image in storage",
	Name:  "user in storage",
	CommonModelFields: models.CommonModelFields{
		ID: 22,
	},
}
