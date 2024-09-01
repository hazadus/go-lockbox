format:
	go fmt ./...
build:
	go build -o ./bin/lockbox ./cmd/lockbox/main.go
test:
	go test -v ./...
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
