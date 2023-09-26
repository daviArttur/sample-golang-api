package usecase

import (
	"testing"

	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	"github.com/daviArttur/sample-golang-api/internal/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	// Stub
	email := "test@mail"
	password := "test_password"

	t.Run("success", func(t *testing.T) {
		// Mock
		repoMock := new(mock.UserRepositoryMock)

		// Arrange
		expectedUser := entity.CreateUser(email, password)
		repoMock.On("FindByEmail", email).Return(nil, nil)
		repoMock.On("Save", *expectedUser).Return(nil)
		usecase := SignUp{UserRepository: repoMock}

		// Act
		usecase.Perform(dto.SignUp{Email: email, Password: password})

		// Assert
		repoMock.AssertCalled(t, "Save", *expectedUser)
	})

	t.Run("error on findByEmail()", func(t *testing.T) {
		// Mock
		repoMock := new(mock.UserRepositoryMock)

		// Arrange
		expectedException := &exception.AppException{Status: 500, Msg: "test"}
		repoMock.On("FindByEmail", email).Return(&entity.User{}, expectedException)
		usecase := SignUp{UserRepository: repoMock}

		// Act
		ex := usecase.Perform(dto.SignUp{Email: email, Password: password})

		// Assert
		assert.Equal(t, expectedException, ex)
	})

	t.Run("user with e-mail already exist", func(t *testing.T) {
		// Mock
		repoMock := new(mock.UserRepositoryMock)

		// Arrange
		repoMock.On("FindByEmail", email).Return(&entity.User{}, nil)
		usecase := SignUp{UserRepository: repoMock}

		// Act
		ex := usecase.Perform(dto.SignUp{Email: email, Password: password})

		// Assert
		assert.Equal(t, *exception.UserAlreadyExist, *ex)
	})
}
