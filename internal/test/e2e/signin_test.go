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
	"github.com/daviArttur/sample-golang-api/internal/app/infra/query"

	e2eContainer "github.com/daviArttur/sample-golang-api/internal/test/config/container"
	e2eContext "github.com/daviArttur/sample-golang-api/internal/test/config/context"
	e2eServer "github.com/daviArttur/sample-golang-api/internal/test/config/server"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // used by migrator
	"github.com/stretchr/testify/assert"

	_ "github.com/golang-migrate/migrate/v4/source/file" // used by migrator
)

type output struct {
	AccessToken string `json:"accessToken"`
}

func TestSignIn(t *testing.T) {
	// Stub
	body := dto.SignIn{Email: "test@mail.com", Password: "1234567"}
	bodyJSON, _ := json.Marshal(body)
	r := bytes.NewReader(bodyJSON)

	// Arrange
	ctx, cancel := e2eContext.GetContext()
	defer cancel()
	connString := e2eContainer.CreatePgTestContainer(ctx)
	app := e2eServer.Run()

	db, err := sql.Open("postgres", connString)

	if err != nil {
		t.Error(err)
	}

	userStub := entity.CreateUser("test@mail.com", "123123")

	db.Exec(query.CreateUser, userStub.Email, userStub.Password)

	// Act
	req, err := http.NewRequest("GET", "/users", r)

	if err != nil {
		t.Error(err)
	}

	defer req.Body.Close()

	response := httptest.NewRecorder()

	app.ServeHTTP(response, req)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	var responseOutput output

	// Assert
	err = json.NewDecoder(response.Body).Decode(&responseOutput)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(responseOutput)
	assert.Equal(t, 200, response.Result().StatusCode)
}
