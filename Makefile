
PHONY: build

build:
	go build -o ./mws ./main.go

cover-html:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

clean:
	@rm -f mws coverage.out