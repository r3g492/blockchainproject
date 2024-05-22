build:
	go build -o bin/blocker

run: build
	./bin/docker

test:
	go test -v ./...