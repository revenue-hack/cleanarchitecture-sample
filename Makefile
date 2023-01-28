.PHONY: generate test

test:
	go test ./...
generate:
	cd src && rm -rf mock && go generate ./...

build:
	go build -o ./sample ./src/cmd/main.go

start:
	./sample
run:
	go run ./src/cmd/main.go
