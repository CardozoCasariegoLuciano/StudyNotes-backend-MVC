package authservice

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	errortypes "CardozoCasariegoLuciano/StudyNotes/helpers/errorTypes"
	"CardozoCasariegoLuciano/StudyNotes/testHelps/mocks/repository"
	"CardozoCasariegoLuciano/StudyNotes/testHelps/mocks/utils"
	logindto "CardozoCasariegoLuciano/StudyNotes/testHelps/objectGenerator/LoginDto"
	registerdto "CardozoCasariegoLuciano/StudyNotes/testHelps/objectGenerator/RegisterDto"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo repository.MockIStorage
var cripto utils.MockIbcrypt
var authServ IAuthService

func TestMain(m *testing.M) {
	repo = repository.MockIStorage{}
	cripto = utils.MockIbcrypt{}
	authServ = NewAuthService(&repo, &cripto)

	//TestRegister
	repo.EXPECT().
		FindUserByEmail(registerdto.RegisterEmailTaken.Email).
		Return(models.User{CommonModelFields: models.CommonModelFields{ID: 20}})

	repo.EXPECT().
		FindUserByEmail(registerdto.RegisterInternalError.Email).
		Return(models.User{})

	repo.EXPECT().
		FindUserByEmail(registerdto.RegisterSucces.Email).
		Return(models.User{})

	cripto.EXPECT().
		HashPassword(registerdto.RegisterInternalError.Password).
		Return("", errors.New(""))

	cripto.EXPECT().
		HashPassword(registerdto.RegisterSucces.Password).
		Return("hased", nil)

	//TestLogin
	repo.EXPECT().
		FindUserByEmail(logindto.LoginEmailDontExist.Email).
		Return(
			models.User{
				CommonModelFields: models.CommonModelFields{ID: 0},
			},
		)

	repo.EXPECT().
		FindUserByEmail(logindto.LoginHasEmailButPassDontMatch.Email).
		Return(
			models.User{
				CommonModelFields: models.CommonModelFields{ID: 12},
			},
		)

	repo.EXPECT().
		FindUserByEmail(logindto.LoginSuccesfuly.Email).
		Return(logindto.UserInStorageToLogin)

	cripto.EXPECT().
		Compare(mock.Anything, logindto.LoginSuccesfuly.Password).
		Return(nil)

	cripto.EXPECT().
		Compare(mock.Anything, logindto.LoginHasEmailButPassDontMatch.Password).
		Return(errors.New(""))

	//both
	repo.EXPECT().
		SaveUser(mock.Anything).
		Return(nil)

	code := m.Run()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	testCases := []struct {
		title                 string
		expectedErrorResponse error
		expectedUserResponse  *responseDto.UserDto
		param                 requestDto.RegisterUserDto
	}{
		{
			title:                 "Mail already taken",
			param:                 registerdto.RegisterEmailTaken,
			expectedErrorResponse: errortypes.MailAlreadyTaken,
			expectedUserResponse:  nil,
		},
		{
			title:                 "Error hashing the password",
			param:                 registerdto.RegisterInternalError,
			expectedErrorResponse: errortypes.InternalError,
			expectedUserResponse:  nil,
		},
		{
			title:                 "Successfully Register",
			param:                 registerdto.RegisterSucces,
			expectedErrorResponse: nil,
			expectedUserResponse: &responseDto.UserDto{
				Role:  registerdto.UserRegistered.Role,
				Name:  registerdto.UserRegistered.Name,
				Email: registerdto.UserRegistered.Email,
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			//Call method under test
			user, err := authServ.RegisterUser(tc.param)

			assert.Equal(t, tc.expectedUserResponse, user)
			assert.Equal(t, tc.expectedErrorResponse, err)
		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		title         string
		entryData     requestDto.LoginUserDto
		expectedError error
		expectedUser  *responseDto.UserDto
	}{
		{
			title:         "when user has not account",
			entryData:     logindto.LoginEmailDontExist,
			expectedUser:  nil,
			expectedError: errortypes.WrongPassOrEmail,
		},
		{
			title:         "when user has account but passwords dont match",
			entryData:     logindto.LoginHasEmailButPassDontMatch,
			expectedUser:  nil,
			expectedError: errortypes.WrongPassOrEmail,
		},
		{
			title:     "Succesfuly Login",
			entryData: logindto.LoginSuccesfuly,
			expectedUser: &responseDto.UserDto{
				ID:    logindto.UserInStorageToLogin.ID,
				Role:  logindto.UserInStorageToLogin.Role,
				Name:  logindto.UserInStorageToLogin.Name,
				Email: logindto.UserInStorageToLogin.Email,
				Image: logindto.UserInStorageToLogin.Image,
			},
			expectedError: nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()
			//Preapare test

			// Call method under test
			user, err := authServ.LoginUser(tc.entryData)

			// Assertions
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedUser, user)
		})
	}
}
