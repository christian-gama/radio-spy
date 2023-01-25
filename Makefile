NAME = "radio-spy"

.PHONY: test
test:
		@echo "Running tests..."
		@go test -v ./...

.PHONY: build-windows
build-windows:
		@echo "Building for Windows..."
		@GOOS=windows GOARCH=amd64 go build -o ./bin/windows/amd64/$(NAME).exe
 
.PHONY: build-linux
build-linux:
		@echo "Building for Linux..."
		@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd64/$(NAME)
		@chmod +x ./bin/linux/amd64/$(NAME)

.PHONY: build-macos
build-macos:
		@echo "Building for MacOS..."
		@GOOS=darwin GOARCH=arm64 go build -o ./bin/macos/arm64/$(NAME)
		@chmod +x ./bin/macos/arm64/$(NAME)

.PHONY: build
build: build-windows build-linux build-macos

.PHONY: clean
clean:
		@echo "Cleaning..."
		@rm -rf ./bin

.PHONY: run
run:
		@echo "Running..."
		@go run main.go