package middleware

import (
	//"go/token"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	UserId int
	Name   string
	jwt.RegisteredClaims
}

func GenerateTokenJwt(UserId int, Name string) string {
	var claims = jwtCustomClaims{
		UserId: UserId,
		Name:   Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	resultJWT, _ := token.SignedString([]byte("123"))
	return resultJWT

}
