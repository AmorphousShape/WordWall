MAIN_PACKAGE_PATH := 
BINARY_NAME = 

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## build: build the application
build:
	go build -o=${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
run: build
	./${BINARY_NAME}

## clean: cleans up object files and removes binaries
clean:
	go clean
	rm ${BINARY_NAME}

## test: run all tests
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## tidy: format code and tidy modfile
tidy:
	go fmt ./...
	go mod tidy -v

## doc: generate a godocs web page running at localhost6060
doc:
	godoc -http=localhost:6060 
