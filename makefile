compile: build run

build:
	go build  -o bin/exec cmd/main.go

run:
	./bin/exec
