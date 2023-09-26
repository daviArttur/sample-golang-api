package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daviArttur/sample-golang-api/internal/app/application/dto"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	e2eContainer "github.com/daviArttur/sample-golang-api/internal/test/config/container"
	e2eContext "github.com/daviArttur/sample-golang-api/internal/test/config/context"
	e2eServer "github.com/daviArttur/sample-golang-api/internal/test/config/server"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // used by migrator
	_ "github.com/golang-migrate/migrate/v4/source/file"       // used by migrator
	"github.com/stretchr/testify/assert"
)

type Dto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestSignUp(t *testing.T) {
	// Arrange
	ctx, cancel := e2eContext.GetContext()
	defer cancel()
	connString := e2eContainer.CreatePgTestContainer(ctx)
	app := e2eServer.Run()

	body := dto.SignUp{Email: "test@mail.com", Password: "1234567"}
	bodyJSON, _ := json.Marshal(body)
	r := bytes.NewReader(bodyJSON)

	// Act
	req, err := http.NewRequest("POST", "/users", r)

	if err != nil {
		t.Error(err)
	}
	defer req.Body.Close()
	response := httptest.NewRecorder()
	app.ServeHTTP(response, req)

	// Assert
	db, err := sql.Open("postgres", connString)

	if err != nil {
		t.Error(err)
	}

	row := db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", body.Email)

	if err != nil {
		t.Error(err)
	}

	var user entity.User

	err = row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		t.Fatal(err)
	}

	row.Err()

	// Assert
	assert.Equal(t, 201, response.Result().StatusCode)
	assert.Equal(t, body.Email, user.Email)
	assert.Equal(t, body.Password, user.Password)
}
