package usecase

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/contract"
	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
)

type ISignIn interface {
	Perform(dto.SignIn) (*SignInOutPut, *exception.AppException)
}

type SignIn struct {
	UserRepository contract.UserRepository
	Token          contract.Jwt
}

func (s *SignIn) Perform(d dto.SignIn) (*SignInOutPut, *exception.AppException) {

	user, ex := s.UserRepository.FindByEmail(d.Email)

	if ex != nil {
		return nil, ex
	}

	if user.Email != d.Email {
		return nil, &exception.AppException{Status: 409, Msg: "asd"}
	}

	token, ex := s.Token.Sign(*user)

	if ex != nil {
		return nil, ex
	}

	return &SignInOutPut{AccessToken: token}, nil
}
