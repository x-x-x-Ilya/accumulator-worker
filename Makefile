upgrade:
	go get -u ./... && go mod tidy

generate:
	go generate ./...

lint:
	 golangci-lint run ./...

test:
	go test -v ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html
	open coverage.html