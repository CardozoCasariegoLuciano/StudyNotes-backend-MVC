package middlewares

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	testhelpers "CardozoCasariegoLuciano/StudyNotes/testHelps/Testhelpers"
	"CardozoCasariegoLuciano/StudyNotes/testHelps/mocks/utils"
	middlewaresdto "CardozoCasariegoLuciano/StudyNotes/testHelps/objectGenerator/MiddlewaresDto"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var tokenMock utils.MockItokens
var middleSUT Imiddlewares

var invalidToken = "invalidToken"
var validToken = "validToken"

func TestMain(m *testing.M) {
	tokenMock = utils.MockItokens{}
	middleSUT = NewMidldeware(&tokenMock)

	tokenMock.EXPECT().ParseToken(invalidToken).Return(nil, errors.New(""))
	tokenMock.EXPECT().ParseToken(validToken).Return(&middlewaresdto.ValidJwtDto, nil)

	code := m.Run()
	os.Exit(code)
}

func TestValidateTokenMiddleware_BadCases(t *testing.T) {
	testCases := []struct {
		title        string
		token        string
		expectedCode int
		expectedData responseDto.ResponseDto
	}{
		{
			title:        "No token sended",
			token:        "",
			expectedCode: http.StatusUnauthorized,
			expectedData: responseDto.ResponseDto{
				Data:        nil,
				MessageType: errorcodes.NO_TOKEN,
				Message:     "Dont have a token",
			},
		},

		{
			title:        "invalid token sended",
			token:        invalidToken,
			expectedCode: http.StatusUnauthorized,
			expectedData: responseDto.ResponseDto{
				Data:        nil,
				MessageType: errorcodes.WRONG_TOKEN,
				Message:     "Wrong or invalid token",
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		testConfig := testhelpers.InitTestConfig{Token: tc.token}

		testData := testhelpers.SetGenericTestData(&testConfig)
		context := *testData.Context
		writer := testData.Recoder

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			//Call method under test
			test := middle.ValidateToken(func(c echo.Context) error {
				return nil
			})
			err := test(context)

			resp := responseDto.NewResponse("", "", nil)
			err = json.Unmarshal(writer.Body.Bytes(), &resp)
			assert.NoError(t, err)

			// Test Cases
			assert.Equal(t, tc.expectedCode, writer.Code)
			assert.Equal(t, tc.expectedData.Data, resp.Data)
			assert.Equal(t, tc.expectedData.MessageType, resp.MessageType)
			assert.Equal(t, tc.expectedData.Message, resp.Message)
		})
	}
}

func TestValidateTokenMiddleware_GoodCases(t *testing.T) {
	testCases := []struct {
		title                     string
		token                     string
		expected_ID_in_Context    uint
		expected_Email_in_Context string
		expected_Role_in_Context  string
	}{
		{
			title:                     "good token sended",
			token:                     validToken,
			expected_ID_in_Context:    middlewaresdto.ValidJwtDto.Id,
			expected_Email_in_Context: middlewaresdto.ValidJwtDto.Email,
			expected_Role_in_Context:  middlewaresdto.ValidJwtDto.Role,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		testConfig := testhelpers.InitTestConfig{Token: tc.token}

		testData := testhelpers.SetGenericTestData(&testConfig)
		context := *testData.Context

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			//Call method under test
			test := middle.ValidateToken(func(c echo.Context) error {
				return nil
			})
			test(context)

			// Test Cases
			assert.Equal(t, tc.expected_Email_in_Context, context.Get("userEmail"))
			assert.Equal(t, tc.expected_Role_in_Context, context.Get("userRole"))
			assert.Equal(t, tc.expected_ID_in_Context, context.Get("userID"))
		})
	}
}
