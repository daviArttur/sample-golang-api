package main

import (
	"database/sql"
	"time"

	"github.com/daviArttur/sample-golang-api/internal/app/infra/router"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID       sql.NullInt64
	Email    string
	Password time.Time
}

func main() {

	g := gin.Default()

	router.SetUp(g)

	g.Run("127.0.0.1:8080")
}
