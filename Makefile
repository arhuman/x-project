BINARY_NAME=api

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

## build: build the binary (${BINARY_NAME})
build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} cmd/api/api.go
#  GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
#  GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

## clean: clean go artefacts (binary included)
clean:
	go clean
	rm ${BINARY_NAME}
#	rm ${BINARY_NAME}-darwin
#	rm ${BINARY_NAME}-windows

## compose_build: docker-compose build
compose_build:
	docker-compose build

## compose_run: stop/build and relaunch docker-compose
compose_run: compose_stop compose_build
	docker-compose up -d

## compose_stop: docker-compose down
compose_stop:
	biomassx && docker-compose down

## doc: makes documentation
.PHONY: doc
doc:
	swag init -g cmd/api/api.go

## local: launch docker-compose env
local: build compose_run

## release: test build and audit current code
release: test build audit

## build: build the binary
run: build
	./${BINARY_NAME}

## test: launch quick tests
test: 
	go test -v ./...

## fulltest: launch full test on docker-compose environment
fulltest: compose_run
	BIOMASSX_FULLTEST=true DBHOST=localhost DBPORT=23306 go test -v ./...

## help: display this usage
.PHONY: help
help:
	@echo 'Usage:'
	@echo ${MAKEFILE_LIST}
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
