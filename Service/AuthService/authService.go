package authservice

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	mysql "CardozoCasariegoLuciano/StudyNotes/Repository/MySql"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	"CardozoCasariegoLuciano/StudyNotes/helpers/roles"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"
	"fmt"
	"net/http"
	"sync"

	"github.com/devfeel/mapper"
	"golang.org/x/crypto/bcrypt"
)

var authS *authService
var once sync.Once

type authService struct {
	storage repository.IStorage
}

func NewAuthService() *authService {
	once.Do(func() {
		fmt.Println("Pasa por aca authService dentro del once")
		authS = &authService{storage: mysql.NewDataBase()}
	})
	return authS
}

func (auth *authService) RegisterUser(user requestDto.RegisterUserDto) (responseDto.ResponseDto, int) {
	//Validate email
	userEmail := auth.storage.FindUserByEmail(user.Email)
	if userEmail.ID != 0 {
		resp := responseDto.NewResponse(
			errorcodes.MAIL_TAKEN,
			"El email ya ha sido tomado",
			nil,
		)
		return resp, http.StatusBadRequest
	}

	//Hashing ths password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response := responseDto.NewResponse(
			errorcodes.HASH_PASS_ERROR,
			"Error hashing the password",
			nil,
		)
		return response, http.StatusInternalServerError
	}

	userM := models.User{Role: roles.USER}
	mapper.AutoMapper(&user, &userM)

	//Save user
	userM.Password = string(hashedPass)
	err = auth.storage.SaveUser(&userM)
	if err != nil {
		response := responseDto.NewResponse(errorcodes.FAIL_SAVING, "trouble saving the user", nil)
		return response, http.StatusInternalServerError
	}

	//Generate Token
	t, err := utils.GenerateToken(userM)
	if err != nil {
		response := responseDto.NewResponse(errorcodes.JWT_ERROR, "trouble creating a JWT", nil)
		return response, http.StatusInternalServerError
	}

	userDto := responseDto.UserDto{ID: int(userM.CommonModelFields.ID)}
	mapper.AutoMapper(&userM, &userDto)
	userToken := responseDto.UserTokenDto{User: userDto, Token: t}

	resp := responseDto.NewResponse("OK", "Usuario creado", userToken)
	return resp, http.StatusOK
}
