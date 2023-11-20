package middlewares

import (
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"
	"sync"

	"github.com/labstack/echo/v4"
)

type Imiddlewares interface {
	ValidateToken(next echo.HandlerFunc) echo.HandlerFunc
}

type Middlewares struct {
	tokenService utils.Itokens
}

var middle *Middlewares
var once sync.Once

func NewMidldeware(tokenService utils.Itokens) Imiddlewares {
	once.Do(func() {
		middle = &Middlewares{
			tokenService: tokenService,
		}
	})
	return middle
}
