package testhelpers

import (
	customvalidator "CardozoCasariegoLuciano/StudyNotes/helpers/customValidator"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

type InitTestConfig struct {
	Path    string
	Method  string
	ReqBody interface{}
	Token   string
}

type TestData struct {
	Echo    *echo.Echo
	Request *http.Request
	Recoder *httptest.ResponseRecorder
	Context *echo.Context
}

type requestData struct {
	body   io.Reader
	path   string
	method string
}

func SetGenericTestData(config *InitTestConfig) *TestData {
	//Init objects
	e := echo.New()
	e.Validator = customvalidator.NewCustomValidator()

	//Create new request, recorder(writer) and contetx
	request := httptest.NewRequest(
		prepareRequest(config).method,
		prepareRequest(config).path,
		prepareRequest(config).body,
	)

	request.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()
	context := e.NewContext(request, writer)

	if len(config.Token) > 0 {
		cookie := http.Cookie{
			Name:  utils.CookieName,
			Value: config.Token,
			Path:  "/",
		}
		context.Request().AddCookie(&cookie)
	}

	returnValues := &TestData{
		Echo:    e,
		Request: request,
		Context: &context,
		Recoder: writer,
	}

	return returnValues
}

func prepareRequest(data *InitTestConfig) requestData {
	resp := requestData{}

	if data.Method == "" {
		resp.method = http.MethodGet
	} else {
		resp.method = data.Method
	}

	if data.Path == "" {
		resp.path = "/"
	} else {
		resp.path = data.Path
	}

	if data.ReqBody == nil {
		resp.body = nil
	} else {
		body, err := json.Marshal(data.ReqBody)
		if err != nil {
			log.Println("Error at marshal ReqBody")
		}
		resp.body = strings.NewReader(string(body))
	}
	return resp
}
