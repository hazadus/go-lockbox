build:
	go build -o lockbox ./cmd/lockbox/main.go
test:
	go test -v
	go test -v ./cmd/lockbox