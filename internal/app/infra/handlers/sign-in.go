package handlers

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SignInHandler struct {
	SignIn usecase.ISignIn
}

func (u *SignInHandler) Handle(c *gin.Context) {

	var dto dto.SignIn

	c.ShouldBindJSON(&dto)
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(dto)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	output, ex := u.SignIn.Perform(dto)

	if ex != nil {
		c.JSON(ex.Status, ex.Msg)
		return
	}

	c.JSON(200, output)
}
