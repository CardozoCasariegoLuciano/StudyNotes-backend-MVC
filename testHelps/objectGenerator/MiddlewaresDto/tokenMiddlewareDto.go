package middlewaresdto

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	"CardozoCasariegoLuciano/StudyNotes/helpers/roles"
)

var ValidJwtDto responseDto.JwtDto = responseDto.JwtDto{
	Email: "validInJWT@emial.com",
	Role:  roles.SUPER_ADMIN,
	Id:    33,
}
