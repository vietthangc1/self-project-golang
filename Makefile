di: 
	wire ./cmd/server

run:
	go run ./cmd/server

lint:
	golangci-lint run --timeout 10m
