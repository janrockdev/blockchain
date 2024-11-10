build:
	go build -o ./bin/node

run: build
	./bin/node

test:
	go test -v ./...