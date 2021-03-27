git:
	git add .
	git commit -m "$(msg)"
	git push origin master

fmt:
	@echo "formatting code"
	go fmt ./...

lint:
	@echo "Linting source code"
	golint ./...

vet:
	@echo "Checking for code issues"
	go vet ./...

test:
	@echo "running tests"
	go test ./...

install:
	@echo "installing external dependencies"
	go mod download

graphql:
	@echo "generating graphql stubs"
	go run github.com/99designs/gqlgen generate

createdb:
	docker exec -it reddit-postgres createdb --username=root --owner=root redditdb

dropdb:
	docker exec -it reddit-postgres dropdb redditdb

migrationup:
	migrate -path ./pkg/db/migration -database "postgres://root:secret@localhost:5432/redditdb?sslmode=disable" -verbose up

migrationdown:
	migrate -path ./pkg/db/migration -database "postgres://root:secret@localhost:5432/redditdb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

run:
	 go run cmd/main.go

dev:
	docker-compose -f dev.yml up -d