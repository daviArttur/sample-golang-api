package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SignInUseCaseMock struct {
	mock.Mock
}

func (u *SignInUseCaseMock) Perform(d dto.SignIn) (*usecase.SignInOutPut, *exception.AppException) {

	args := u.Called(d)

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	if args.Get(1) == nil {
		return args.Get(0).(*usecase.SignInOutPut), nil
	}

	return args.Get(0).(*usecase.SignInOutPut), args.Get(1).(*exception.AppException)
}

func TestSignInHandler(t *testing.T) {
	// Stub
	email := "test@mail.com"
	password := "test"

	t.Run("success", func(t *testing.T) {
		// Stub
		dtoStub := dto.SignIn{Email: email, Password: password}
		bodyJson, _ := json.Marshal(dtoStub)

		// Mock
		usecaseMock := new(SignInUseCaseMock)
		handler := SignInHandler{SignIn: usecaseMock}

		// Arrange
		expectedOutPut := &usecase.SignInOutPut{AccessToken: "test"}
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		bodyReader := strings.NewReader(string(bodyJson))
		usecaseMock.On("Perform", dtoStub).Return(expectedOutPut, nil)
		c.Request = &http.Request{
			Body: io.NopCloser(bodyReader),
		}

		// Act
		handler.Handle(c)
		text, _ := io.ReadAll(recorder.Result().Body)
		wantBody := string(text)
		wantStatusCode := 200
		gotBody := `{"accessToken":"test"}`
		gotStatusCode := recorder.Result().StatusCode

		// Assert
		assert.Equal(t, wantStatusCode, gotStatusCode)
		assert.Equal(t, wantBody, gotBody)
	})

	t.Run("error on validate input", func(t *testing.T) {
		// Mock
		usecaseMock := new(SignInUseCaseMock)
		handler := SignInHandler{SignIn: usecaseMock}
		var dto dto.SignIn

		// Arrange
		validate := validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(dto) // simulate validation in handler
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder) // not send request body

		// Act
		handler.Handle(c)
		wantBody, _ := json.Marshal(err.Error())
		wantStatusCode := 400
		gotBody, _ := io.ReadAll(recorder.Result().Body)
		gotStatusCode := recorder.Result().StatusCode

		// Assert
		assert.Equal(t, wantStatusCode, gotStatusCode)
		assert.Equal(t, wantBody, gotBody)
	})

	t.Run("error on usecase", func(t *testing.T) {
		// Stub
		dtoStub := dto.SignIn{Email: email, Password: password}
		bodyJson, _ := json.Marshal(dtoStub)

		// Mock
		usecaseMock := new(SignInUseCaseMock)
		handler := SignInHandler{SignIn: usecaseMock}

		// Arrange
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		bodyReader := strings.NewReader(string(bodyJson))
		usecaseMock.On("Perform", dtoStub).Return(&usecase.SignInOutPut{}, &exception.AppException{Status: 500, Msg: "test"})
		c.Request = &http.Request{
			Body: io.NopCloser(bodyReader),
		}

		// Act
		handler.Handle(c)
		wantBody, _ := json.Marshal("test")
		wantStatusCode := 500
		gotBody, _ := io.ReadAll(recorder.Result().Body)
		gotStatusCode := recorder.Result().StatusCode

		// Assert
		assert.Equal(t, wantStatusCode, gotStatusCode)
		assert.Equal(t, wantBody, gotBody)
	})
}
