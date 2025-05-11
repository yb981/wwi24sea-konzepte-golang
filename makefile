# all
all:	build test

# build all
build:
	go build ./...

# test all packages
test:
	go test ./...

# coverage
coverage:
	go test ./... -cover

coverage-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

coverage-report:
	go tool cover -func=coverage.out

clean:
	rm coverage.out coverage.html