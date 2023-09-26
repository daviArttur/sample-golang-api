package modules

import (
	"database/sql"

	"github.com/daviArttur/sample-golang-api/internal/app/application/usecase"
	"github.com/daviArttur/sample-golang-api/internal/app/infra/handlers"
	"github.com/daviArttur/sample-golang-api/internal/app/infra/repository"
	"github.com/daviArttur/sample-golang-api/internal/app/infra/services/token"
)

type Handlers struct {
	SignInHandler handlers.SignInHandler
	SignUpHandler handlers.SignUpHandler
}

type Modules struct {
	DB *sql.DB
}

func (m *Modules) UserModule() Handlers {

	// Repository
	repo := repository.UserRepository{DB: m.DB}

	// Use Cases
	signInUseCase := usecase.SignIn{
		Token:          &token.Jwt{Secret: "12345"},
		UserRepository: repo,
	}

	signInHandler := handlers.SignInHandler{
		SignIn: &signInUseCase,
	}

	signUpUseCase := usecase.SignUp{
		UserRepository: repo,
	}

	signUpHandler := handlers.SignUpHandler{
		Usecase: &signUpUseCase,
	}

	return Handlers{
		SignInHandler: signInHandler,
		SignUpHandler: signUpHandler,
	}
}
