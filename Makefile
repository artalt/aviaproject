GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=./cmd/app/main.out

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: gen-api build

build:
	go build -o ${BINARY_NAME} ./cmd/app/main.go

gen-api:
	go generate ./...

lint:
	golangci-lint run

test:
	go test -v ./cmd/app/main.go

run:
	go build -o ${BINARY_NAME} ./cmd/app/main.go
	./${BINARY_NAME} -c ./cmd/app/config.yaml

clean:
	go clean
	rm ${BINARY_NAME}

migrate-up:
	migrate -source file://migrations/postgresql -database "postgres://root:pass@localhost:5433/digdb?sslmode=disable" up

migrate-down:
	migrate -source file://migrations/postgresql -database "postgres://root:pass@localhost:5433/digdb?sslmode=disable" down

migrate-down-last:
	migrate -source file://migrations/postgresql -database "postgres://root:pass@localhost:5433/digdb?sslmode=disable" down 1
