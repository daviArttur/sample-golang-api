package mock

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	tMock "github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	tMock.Mock
}

func (u *UserRepositoryMock) FindByEmail(email string) (*entity.User, *exception.AppException) {
	args := u.Called(email)

	if args.Get(1) == nil && args.Get(0) == nil {
		return nil, nil
	}

	if args.Get(1) == nil {
		return args.Get(0).(*entity.User), nil
	}

	return args.Get(0).(*entity.User), args.Get(1).(*exception.AppException)
}

func (u *UserRepositoryMock) FindById(id int) (*entity.User, *exception.AppException) {
	args := u.Called(id)

	return args.Get(0).(*entity.User), args.Get(1).(*exception.AppException)
}

func (u *UserRepositoryMock) Save(user entity.User) *exception.AppException {
	args := u.Called(user)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*exception.AppException)
}
