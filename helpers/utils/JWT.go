package utils

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	"CardozoCasariegoLuciano/StudyNotes/configuration"

	"github.com/golang-jwt/jwt"
)

type Itokens interface {
	GenerateToken(user responseDto.UserDto) (string, error)
}

type Token struct{}

func (tk *Token) GenerateToken(user responseDto.UserDto) (string, error) {
	claims := responseDto.JwtDto{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config := configuration.GetConfig()

	return token.SignedString([]byte(config.Jwt.Secret))
}
