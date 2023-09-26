package usecase

import (
	"testing"

	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	"github.com/daviArttur/sample-golang-api/internal/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	// Stub
	email := "test@mail"
	password := "test_password"

	// Mock
	tokenMock := "token"

	t.Run("success", func(t *testing.T) {
		// Stub
		userStub := entity.CreateUser(email, password)

		// Mock
		repoMock := new(mock.UserRepositoryMock)
		jwtMock := new(mock.JwtMock)

		// Arrange
		repoMock.On("FindByEmail", email).Return(userStub, nil)
		jwtMock.On("Sign", *userStub).Return(tokenMock, nil)
		usecase := SignIn{UserRepository: repoMock, Token: jwtMock}

		// Act
		output, ex := usecase.Perform(dto.SignIn{Email: email, Password: password})

		// Assert
		assert.Nil(t, ex)
		assert.Equal(t, SignInOutPut{AccessToken: tokenMock}, *output)
	})

	t.Run("error on findByEmail()", func(t *testing.T) {
		// Stub
		userStub := entity.CreateUser(email, password)

		// Mock
		repoMock := new(mock.UserRepositoryMock)
		jwtMock := new(mock.JwtMock)

		// Arrange
		expectedException := &exception.AppException{Status: 400, Msg: "test"}
		repoMock.On("FindByEmail", email).Return(userStub, expectedException)
		jwtMock.On("Sign", *userStub).Return(tokenMock, nil)
		usecase := SignIn{UserRepository: repoMock, Token: jwtMock}

		// Act
		_, ex := usecase.Perform(dto.SignIn{Email: email, Password: password})

		// Assert
		assert.Equal(t, *expectedException, *ex)
	})

	t.Run("error on compare passwords", func(t *testing.T) {
		// Stub
		diffEmail := "test2@mail"
		userStub := entity.CreateUser(diffEmail, password)

		// Mock
		repoMock := new(mock.UserRepositoryMock)
		jwtMock := new(mock.JwtMock)

		// Arrange
		expectedException := &exception.AppException{Status: 409, Msg: "asd"}
		repoMock.On("FindByEmail", email).Return(userStub, nil)
		jwtMock.On("Sign", *userStub).Return(tokenMock, nil)
		usecase := SignIn{UserRepository: repoMock, Token: jwtMock}

		// Act
		_, ex := usecase.Perform(dto.SignIn{Email: email, Password: password})

		// Assert
		assert.Equal(t, *expectedException, *ex)
	})

	t.Run("error on sign token", func(t *testing.T) {
		// Stub
		userStub := entity.CreateUser(email, password)

		// Mock
		repoMock := new(mock.UserRepositoryMock)
		jwtMock := new(mock.JwtMock)

		// Arrange
		expectedException := &exception.AppException{Status: 409, Msg: "error on sign token"}
		repoMock.On("FindByEmail", email).Return(userStub, nil)
		jwtMock.On("Sign", *userStub).Return(tokenMock, expectedException)
		usecase := SignIn{UserRepository: repoMock, Token: jwtMock}

		// Act
		_, ex := usecase.Perform(dto.SignIn{Email: email, Password: password})

		// Assert
		assert.Equal(t, *expectedException, *ex)
	})
}
