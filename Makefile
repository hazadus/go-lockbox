build:
	go build -o ./bin/lockbox ./cmd/lockbox/main.go
test:
	go test -v ./...
format:
	go fmt ./...
