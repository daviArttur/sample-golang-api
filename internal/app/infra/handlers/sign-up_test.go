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
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SignUpUseCaseMock struct {
	mock.Mock
}

func (u *SignUpUseCaseMock) Perform(d dto.SignUp) *exception.AppException {

	args := u.Called(d)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*exception.AppException)
}

func TestSignUpHandler(t *testing.T) {
	// Stub
	email := "test@mail.com"
	password := "test"

	t.Run("success", func(t *testing.T) {
		// Stub
		dtoStub := dto.SignUp{Email: email, Password: password}
		bodyJson, _ := json.Marshal(dtoStub)

		// Mock
		usecaseMock := new(SignUpUseCaseMock)
		handler := SignUpHandler{Usecase: usecaseMock}

		// Arrange

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		bodyReader := strings.NewReader(string(bodyJson))
		usecaseMock.On("Perform", dtoStub).Return(nil)
		c.Request = &http.Request{
			Body: io.NopCloser(bodyReader),
		}

		// Act
		handler.Handle(c)
		wantStatusCode := 201
		gotStatusCode := recorder.Result().StatusCode

		// Assert
		assert.Equal(t, wantStatusCode, gotStatusCode)
	})

	t.Run("error on validate input", func(t *testing.T) {
		// Mock
		usecaseMock := new(SignUpUseCaseMock)
		handler := SignUpHandler{Usecase: usecaseMock}
		var dto dto.SignUp

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
		dtoStub := dto.SignUp{Email: email, Password: password}
		bodyJson, _ := json.Marshal(dtoStub)

		// Mock
		usecaseMock := new(SignUpUseCaseMock)
		handler := SignUpHandler{Usecase: usecaseMock}

		// Arrange
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		bodyReader := strings.NewReader(string(bodyJson))
		usecaseMock.On("Perform", dtoStub).Return(&exception.AppException{Status: 500, Msg: "test"})
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
