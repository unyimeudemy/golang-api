build:
	@go build -o bin/goland-api cmd/main.go


test:
	@go test -v ./...


run: build
		@./bin/goland-api