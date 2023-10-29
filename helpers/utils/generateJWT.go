package utils

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	"CardozoCasariegoLuciano/StudyNotes/configuration"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user models.User) (string, error) {
	claims := responseDto.JwtDto{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config := configuration.GetConfig()

	return token.SignedString([]byte(config.Jwt.Secret))
}
