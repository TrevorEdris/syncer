BIN=bin
BINARY_NAME=syncer
BINARY_LOCATION=${BIN}/${BINARY_NAME}
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GIT_HASH=$(shell git rev-parse --short HEAD)

show:
	@echo ${GOOS}
	@echo ${GOARCH}
	@echo ${GIT_HASH}

create-bin:
	mkdir -p ./${BIN}

build-linux: create-bin
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_LOCATION}-linux main.go

build-darwin: create-bin
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_LOCATION}-darwin main.go

build-windows: create-bin
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_LOCATION}-windows main.go

package: build-linux build-darwin build-windows
	zip -r ${BINARY_NAME}-${GIT_HASH}.zip ${BINARY_LOCATION}-windows ${BINARY_LOCATION}-darwin ${BINARY_LOCATION}-linux

build: create-bin
	go build -o ${BINARY_LOCATION} main.go

clean:
	go clean
	rm -rf ${BIN}
	rm -rf ${BINARY_NAME}*.zip

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all
