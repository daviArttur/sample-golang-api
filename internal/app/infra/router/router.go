package router

import (
	"github.com/daviArttur/sample-golang-api/internal/app/infra/config/db"
	"github.com/daviArttur/sample-golang-api/internal/app/infra/modules"
	"github.com/gin-gonic/gin"
)

func SetUp(g *gin.Engine) {
	connection := db.CreatePostgresConnection()
	modules := modules.Modules{DB: connection}
	userM := modules.UserModule()
	g.GET("/users", userM.SignInHandler.Handle)
	g.POST("/users", userM.SignUpHandler.Handle)
}
