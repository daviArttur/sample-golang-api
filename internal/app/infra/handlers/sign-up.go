package handlers

import (
	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SignUpHandler struct {
	Usecase usecase.ISignUp
}

func (u *SignUpHandler) Handle(c *gin.Context) {

	var d dto.SignUp

	c.ShouldBindJSON(&d)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(d)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	ex := u.Usecase.Perform(d)

	if ex != nil {
		c.JSON(ex.Status, ex.Msg)
		return
	}

	c.Status(201)
	c.Writer.WriteHeaderNow()
}
