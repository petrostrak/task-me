BINARY_NAME=taskME
APP_NAME=taskME
VERSION=1.0.0

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	rm -f fyne-md
	fyne package -appVersion ${VERSION} -name ${APP_NAME} -release

## run: builds and runs the application
run:
	env DB_PATH="./sql.db" go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...