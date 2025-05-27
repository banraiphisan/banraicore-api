BINARY_NAME=banraicore-api
BUILD_DIR=bin
SRC_DIR=./cmd/

test:
	go test -v -cover -covermode=atomic ./internal/usecase/...

unittest:
	go test -short  ./internal/usecase/...

coverage:
	go test ./internal/usecase/.../service/... ./internal/domain/... -coverprofile=coverage.out

lint:
	golangci-lint run

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

clean:
	rm -rf $(BUILD_DIR)

help:
	@echo "Makefile Targets:"
	@echo "  test        Run tests with coverage"
	@echo "  unittest    Run unit tests in short mode"
	@echo "  lint        Run code linter"
	@echo "  build       Build the binary"
	@echo "  clean       Remove build artifacts"