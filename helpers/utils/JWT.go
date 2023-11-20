package utils

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	"CardozoCasariegoLuciano/StudyNotes/configuration"
	"sync"

	"github.com/golang-jwt/jwt"
)

type Itokens interface {
	GenerateToken(user responseDto.UserDto) (string, error)
	ParseToken(token string) (*responseDto.JwtDto, error)
}

type Token struct{}

var tk *Token
var onceTK sync.Once

func NewToken() Itokens {
	onceTK.Do(func() {
		tk = &Token{}
	})
	return tk
}

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

func (tk *Token) ParseToken(token string) (*responseDto.JwtDto, error) {
	dataToken := responseDto.JwtDto{}
	secret := configuration.GetConfig().Jwt.Secret

	tkn, err := jwt.ParseWithClaims(token, &dataToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !tkn.Valid {
		return nil, err
	}

	return &dataToken, nil
}
