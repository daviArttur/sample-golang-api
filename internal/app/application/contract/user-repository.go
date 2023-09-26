package contract

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, *exception.AppException)
	FindById(id int) (*entity.User, *exception.AppException)
	Save(entity.User) *exception.AppException
}
