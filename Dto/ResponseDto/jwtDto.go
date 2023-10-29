package responseDto

import "github.com/golang-jwt/jwt"

type JwtDto struct {
	Id    int
	Email string
	Role  string
	jwt.StandardClaims
}
