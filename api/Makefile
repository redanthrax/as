BINARY_NAME=api
build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} cmd/api/main.go
tests:
	go test ./...
run:
	go run cmd/api/main.go
func:
	func start --verbose
