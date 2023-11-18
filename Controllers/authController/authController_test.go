package authcontroller

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	errortypes "CardozoCasariegoLuciano/StudyNotes/helpers/errorTypes"
	testhelpers "CardozoCasariegoLuciano/StudyNotes/testHelps/Testhelpers"
	"CardozoCasariegoLuciano/StudyNotes/testHelps/mocks/authservice"
	"CardozoCasariegoLuciano/StudyNotes/testHelps/mocks/utils"
	logindto "CardozoCasariegoLuciano/StudyNotes/testHelps/objectGenerator/LoginDto"
	registerdto "CardozoCasariegoLuciano/StudyNotes/testHelps/objectGenerator/RegisterDto"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var serv authservice.MockIAuthService
var tokenMock utils.MockItokens
var authcontroller *AuthController

func TestMain(m *testing.M) {
	serv = authservice.MockIAuthService{}
	tokenMock = utils.MockItokens{}

	serv.EXPECT().
		RegisterUser(
			registerdto.RegisterEmailTaken,
		).
		Return(
			nil,
			errortypes.MailAlreadyTaken,
		)

	serv.EXPECT().
		RegisterUser(
			registerdto.RegisterInternalError,
		).
		Return(
			nil,
			errortypes.InternalError,
		)

	serv.EXPECT().
		RegisterUser(registerdto.RegisterSucces).
		Return(
			&registerdto.UserRegistered,
			nil,
		)

	serv.EXPECT().
		LoginUser(logindto.LoginEmailDontExist).
		Return(
			nil,
			errortypes.WrongPassOrEmail,
		)

	serv.EXPECT().
		LoginUser(logindto.LoginInternalError).
		Return(
			nil,
			errortypes.InternalError,
		)

	serv.EXPECT().
		LoginUser(logindto.LoginSuccesfuly).
		Return(
			&logindto.UserLoged,
			nil,
		)

	tokenMock.EXPECT().
		GenerateToken(mock.Anything).
		Return("", nil)

	authcontroller = NewAuthController(&serv, &tokenMock)

	code := m.Run()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	testCases := []struct {
		title        string
		expectedCode int
		expectedData responseDto.ResponseDto
		body         interface{}
	}{
		{
			title:        "body types dont match",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Data:        nil,
				MessageType: errorcodes.BODY_TYPES_ERROR,
				Message:     "Error con los datos enviados",
			},
			body: map[string]interface{}{
				"name":         "Rogelio",
				"password":     123123,
				"confirmation": 123123,
				"email":        true,
			},
		},
		{
			title:        "Body required validation errors",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Message:     "Error en la validacion de los datos enviados",
				MessageType: errorcodes.BODY_VALIDATION_ERROR,
				Data: map[string]interface{}{
					"confirmation": "Type error: required",
					"password":     "Type error: required",
					"name":         "Type error: required",
					"email":        "Type error: required",
				},
			},
			body: map[string]string{},
		},
		{
			title:        "Body secondary validations errors",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Message:     "Error en la validacion de los datos enviados",
				MessageType: errorcodes.BODY_VALIDATION_ERROR,
				Data: map[string]interface{}{
					"confirmation": "Type error: eqfield",
					"password":     "Type error: min",
					"name":         "Type error: min",
					"email":        "Type error: email",
				},
			},
			body: map[string]string{
				"name":         "Ro",
				"password":     "12",
				"email":        "testtestcom",
				"confirmation": "123123",
			},
		},
		{
			title:        "Error mail already taken",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Message:     "Email already taken",
				MessageType: errorcodes.MAIL_TAKEN,
				Data:        nil,
			},
			body: map[string]string{
				"name":         registerdto.RegisterEmailTaken.Name,
				"password":     registerdto.RegisterEmailTaken.Password,
				"email":        registerdto.RegisterEmailTaken.Email,
				"confirmation": registerdto.RegisterEmailTaken.Confirmation,
			},
		},
		{
			title:        "Internal error",
			expectedCode: http.StatusInternalServerError,
			expectedData: responseDto.ResponseDto{
				Message:     "Server internal error",
				MessageType: errorcodes.INTERNAL_ERROR,
				Data:        nil,
			},
			body: map[string]string{
				"name":         registerdto.RegisterInternalError.Name,
				"password":     registerdto.RegisterInternalError.Password,
				"email":        registerdto.RegisterInternalError.Email,
				"confirmation": registerdto.RegisterInternalError.Confirmation,
			},
		},
		{
			title:        "Body OK",
			expectedCode: http.StatusCreated,
			expectedData: responseDto.ResponseDto{
				MessageType: "OK",
				Message:     "User created",
				Data: map[string]interface{}{
					"user": map[string]interface{}{
						"name":  registerdto.UserRegistered.Name,
						"role":  registerdto.UserRegistered.Role,
						"image": registerdto.UserRegistered.Image,
						"id":    float64(registerdto.UserRegistered.ID),
						"email": registerdto.UserRegistered.Email,
					}},
			},
			body: map[string]string{
				"name":         registerdto.RegisterSucces.Name,
				"password":     registerdto.RegisterSucces.Password,
				"email":        registerdto.RegisterSucces.Email,
				"confirmation": registerdto.RegisterSucces.Confirmation,
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		testConfig := testhelpers.InitTestConfig{
			ReqBody: tc.body,
		}

		testData := testhelpers.SetGenericTestData(&testConfig)
		context := *testData.Context
		writer := testData.Recoder

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			//Call method under test
			err := authcontroller.Register(context)
			assert.NoError(t, err)

			//Assert
			resp := responseDto.ResponseDto{}
			err = json.Unmarshal(writer.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, writer.Code)
			assert.Equal(t, tc.expectedData.Message, resp.Message)
			assert.Equal(t, tc.expectedData.MessageType, resp.MessageType)
			assert.Equal(t, tc.expectedData.Data, resp.Data)
		})
	}
}

func TestLogIn(t *testing.T) {
	testCases := []struct {
		title        string
		expectedCode int
		expectedData responseDto.ResponseDto
		body         interface{}
	}{
		{
			title:        "body types dont match",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Data:        nil,
				MessageType: errorcodes.BODY_TYPES_ERROR,
				Message:     "Error con los datos enviados",
			},
			body: map[string]interface{}{
				"password": 123123,
				"email":    true,
			},
		},
		{
			title:        "Body required validation errors",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Message:     "Error en la validacion de los datos enviados",
				MessageType: errorcodes.BODY_VALIDATION_ERROR,
				Data: map[string]interface{}{
					"password": "Type error: required",
					"email":    "Type error: required",
				},
			},
			body: map[string]string{},
		},
		{
			title:        "Body secondary validations errors",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Message:     "Error en la validacion de los datos enviados",
				MessageType: errorcodes.BODY_VALIDATION_ERROR,
				Data: map[string]interface{}{
					"email": "Type error: email",
				},
			},
			body: map[string]string{
				"password": "123456",
				"email":    "testtestcom",
			},
		},
		{
			title:        "Error mail dont exist",
			expectedCode: http.StatusBadRequest,
			expectedData: responseDto.ResponseDto{
				Message:     "Wrong email or password",
				MessageType: errorcodes.WRONG_LOGIN_DATA,
				Data:        nil,
			},
			body: map[string]string{
				"password": logindto.LoginEmailDontExist.Password,
				"email":    logindto.LoginEmailDontExist.Email,
			},
		},
		{
			title:        "Internal error",
			expectedCode: http.StatusInternalServerError,
			expectedData: responseDto.ResponseDto{
				Message:     "Server internal error",
				MessageType: errorcodes.INTERNAL_ERROR,
				Data:        nil,
			},
			body: map[string]string{
				"password": logindto.LoginInternalError.Password,
				"email":    logindto.LoginInternalError.Email,
			},
		},
		{
			title:        "Body OK",
			expectedCode: http.StatusOK,
			expectedData: responseDto.ResponseDto{
				MessageType: "OK",
				Message:     "User loged",
				Data: map[string]interface{}{
					"user": map[string]interface{}{
						"name":  logindto.UserLoged.Name,
						"role":  logindto.UserLoged.Role,
						"image": logindto.UserLoged.Image,
						"id":    float64(logindto.UserLoged.ID),
						"email": logindto.UserLoged.Email,
					}},
			},
			body: map[string]string{
				"password": logindto.LoginSuccesfuly.Password,
				"email":    logindto.LoginSuccesfuly.Email,
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		testConfig := testhelpers.InitTestConfig{
			ReqBody: tc.body,
		}

		testData := testhelpers.SetGenericTestData(&testConfig)
		context := *testData.Context
		writer := testData.Recoder

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			//Call method under test
			err := authcontroller.Login(context)
			assert.NoError(t, err)

			//Assert
			resp := responseDto.ResponseDto{}
			err = json.Unmarshal(writer.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, writer.Code)
			assert.Equal(t, tc.expectedData.Message, resp.Message)
			assert.Equal(t, tc.expectedData.MessageType, resp.MessageType)
			assert.Equal(t, tc.expectedData.Data, resp.Data)
		})
	}
}
