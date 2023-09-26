package server

import (
	"github.com/daviArttur/sample-golang-api/internal/app/infra/router"
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	g := gin.Default()

	router.SetUp(g)

	return g
}
