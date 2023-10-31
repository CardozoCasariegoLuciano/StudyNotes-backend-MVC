package responseDto

import "github.com/golang-jwt/jwt"

type JwtDto struct {
	Id    uint
	Email string
	Role  string
	jwt.StandardClaims
}
