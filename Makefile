.PHONY: build
build:
	go build -v ./cmd/apiserver


.DEFAULT_GOAL:=build
 

.PHONY: migrate
migrate:
	migrate -database postgres://localhost/flashcards?sslmode=disable -path migrations up
	migrate -database postgres://localhost/flashcards_test?sslmode=disable -path migrations up


.PHONY: rollback
rollback:
	migrate -database postgres://localhost/flashcards?sslmode=disable -path migrations down
	migrate -database postgres://localhost/flashcards_test?sslmode=disable -path migrations down

.PHONY: test
test:
	go test -v -race -timeout 30s ./...