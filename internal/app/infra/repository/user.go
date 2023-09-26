package repository

import (
	"database/sql"
	"fmt"

	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	"github.com/daviArttur/sample-golang-api/internal/app/infra/query"
)

type UserRepository struct {
	DB *sql.DB
}

func (u UserRepository) FindByEmail(email string) (*entity.User, *exception.AppException) {
	fmt.Println("qwe")

	row := u.DB.QueryRow(query.FindUserByEmail, email)

	var user entity.User

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	fmt.Println("qwe")
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuário não encontrado, sem erro
		}
		return nil, exception.QueryErr // Erro ao escanear valores do banco de dados
	}

	return &user, nil
}

func (u UserRepository) FindById(id int) (*entity.User, *exception.AppException) {
	row := u.DB.QueryRow(query.FindUserById, id)

	var user entity.User

	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Não encontrado
		}
		return nil, exception.QueryErr // Erro inesperado
	}

	return &user, nil // Encontrado
}

func (u UserRepository) Save(user entity.User) *exception.AppException {
	_, err := u.DB.Exec(query.CreateUser, user.Email, user.Password)

	if err != nil {
		return exception.QueryErr // Erro ao criar usuário
	}

	return nil
}
