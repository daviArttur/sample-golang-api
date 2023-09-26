package usecase

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/contract"
	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
)

type ISignUp interface {
	Perform(d dto.SignUp) *exception.AppException
}

type SignUp struct {
	UserRepository contract.UserRepository
}

type SignInOutPut struct {
	AccessToken string `json:"accessToken"`
}

func (s *SignUp) Perform(d dto.SignUp) *exception.AppException {

	user, ex := s.UserRepository.FindByEmail(d.Email)

	if ex != nil {
		return ex
	}

	if user != nil {
		return exception.UserAlreadyExist
	}

	user = entity.CreateUser(d.Email, d.Password)

	ex = s.UserRepository.Save(*user)

	return ex
}
