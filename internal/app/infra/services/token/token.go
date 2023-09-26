package token

import (
	"time"

	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	"github.com/golang-jwt/jwt"
)

type Jwt struct {
	Secret string
}

func (j *Jwt) Sign(user entity.User) (string, *exception.AppException) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(5 * time.Minute),
	})

	tokenString, err := token.SignedString([]byte(j.Secret))

	if err != nil {
		return tokenString, &exception.AppException{Status: 500, Msg: err.Error()}
	}

	return tokenString, nil
}
