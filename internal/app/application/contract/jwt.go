package contract

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
)

type Jwt interface {
	Sign(entity.User) (string, *exception.AppException)
}
