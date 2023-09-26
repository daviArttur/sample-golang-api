build:
    go build -o meu_app main.go

test:
    go test ./...

test-cov:
    go test ./... -cover

run:
    go run cmd/app/main.go

migrate-up:
		migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migrate-down:
		migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down