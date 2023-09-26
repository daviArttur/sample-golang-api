package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/daviArttur/sample-golang-api/internal/app/application/exception"
	"github.com/daviArttur/sample-golang-api/internal/app/domain/entity"
	"github.com/daviArttur/sample-golang-api/internal/app/infra/query"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {

	t.Run("FindByEmail - it should return an user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		rows := sqlmock.NewRows([]string{"id", "email", "passwod"}).AddRow(1, "email", "password")
		userRepository := &UserRepository{DB: db}
		mock.ExpectQuery(`SELECT id, email, password FROM users WHERE email = $1;`).WithArgs("email").WillReturnRows(rows).RowsWillBeClosed()

		expectedUser := entity.User{ID: 1, Email: "email", Password: "password"}

		// Act
		user, ex := userRepository.FindByEmail("email")

		// Assert
		assert.Equal(t, *user, expectedUser)
		assert.Nil(t, ex)
		//assert.Equal(t, exist, true)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal("got error:", err)
		}
	})

	t.Run("FindByEmail - it should return exist param equal to false because user does not exist", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		rows := sqlmock.NewRows([]string{})
		userRepository := &UserRepository{DB: db}
		mock.ExpectQuery(`SELECT id, email, password FROM users WHERE email = $1;`).WithArgs("email").WillReturnRows(rows).RowsWillBeClosed()

		// Act
		user, ex := userRepository.FindByEmail("email")

		// Assert 2
		assert.Nil(t, user)
		assert.Nil(t, ex)
	})

	t.Run("FindByEmail - it should return exist param equal to false because user does not exist", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		userRepository := &UserRepository{DB: db}
		mock.ExpectQuery(`SELECT id, email, password FROM users WHERE email = $1;`).WithArgs("email").WillReturnError(errors.New("")).RowsWillBeClosed()

		// Act
		user, ex := userRepository.FindByEmail("email")

		// Assert 2
		assert.Equal(t, ex, exception.QueryErr)
		assert.Nil(t, user)
	})

	t.Run("Save - it should create an user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		userRepository := &UserRepository{DB: db}

		// Assert 1
		mock.ExpectExec(`INSERT INTO users (email, password) VALUES ($1, $2);`).WithArgs("test", "123").WillReturnResult(sqlmock.NewResult(1, 1))

		// Act
		userRepository.Save(entity.User{ID: 1, Email: "test", Password: "123"})
	})

	t.Run("Save - return exception because an error ocurred on insert user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		userRepository := &UserRepository{DB: db}

		// Act
		mock.ExpectExec(`INSERT INTO users (email, password) VALUES ($1, $2);`).WithArgs("test", "123").WillReturnError(errors.New(""))
		ex := userRepository.Save(entity.User{ID: 1, Email: "test", Password: "123"})

		// Assert
		assert.Equal(t, exception.QueryErr, ex)
	})

	t.Run("FindById - it should return an user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		expectedUser := entity.User{ID: 1, Email: "email", Password: "password"}
		rows := sqlmock.NewRows([]string{"id", "email", "passwod"}).AddRow(1, "email", "password")
		userRepository := &UserRepository{DB: db}
		mock.ExpectQuery(query.FindUserById).WithArgs(1).WillReturnRows(rows).RowsWillBeClosed()

		// Act
		user, ex := userRepository.FindById(1)

		// Assert 2
		assert.Equal(t, expectedUser, *user)
		assert.Nil(t, ex)
	})

	t.Run("FindById - it should return afdn user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		rows := sqlmock.NewRows([]string{"id", "email", "passwod"})
		userRepository := &UserRepository{DB: db}
		mock.ExpectQuery(query.FindUserById).WithArgs(1).WillReturnRows(rows).RowsWillBeClosed()

		// Act
		user, ex := userRepository.FindById(1)

		// Assert 2
		assert.Nil(t, ex)
		assert.Nil(t, user)
	})

	t.Run("FindById - it should return an user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// Arrange
		userRepository := &UserRepository{DB: db}
		mock.ExpectQuery(query.FindUserById).WithArgs(1).WillReturnError(errors.New("")).RowsWillBeClosed()

		// Act
		user, ex := userRepository.FindById(1)

		// Assert 2
		assert.Equal(t, ex, exception.QueryErr)
		assert.Nil(t, user)
	})
}
