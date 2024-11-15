build:
	go build -o ./bin/node

run: build
	./bin/node

test:
	go test ./...

testv:
	go test -v ./...