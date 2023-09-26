package test_container

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var dbName = "users"
var dbUser = "user"
var dbPassword = "password"

func CreatePgTestContainer(ctx context.Context) string {

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := postgresContainer.Terminate(ctx); err != nil {
	// 		fmt.pi
	// 		panic(err)
	// 	}
	// }()

	urlConnection, err := pgContainer.ConnectionString(ctx)

	if err != nil {
		panic(err)
	}

	connString := urlConnection + "sslmode=disable"

	runMigrate(connString)

	err = os.Setenv("DB_URL", connString)

	if err != nil {
		panic(err)
	}

	return connString
}

func runMigrate(connString string) {
	migration, err := migrate.New("file:C:\\Users\\davia\\Desktop\\code\\sample-golang-api\\migrations", connString)

	if err != nil {
		panic(err)
	}

	defer migration.Close()

	err = migration.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
