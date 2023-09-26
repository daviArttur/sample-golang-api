package mock

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	tMock "github.com/stretchr/testify/mock"
)

type JwtMock struct {
	tMock.Mock
}

func (j *JwtMock) Sign(u entity.User) (string, *exception.AppException) {

	args := j.Called(u)

	if args.Get(1) == nil {
		return args.String(0), nil
	}

	return args.String(0), args.Get(1).(*exception.AppException)
}
