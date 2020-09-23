.PHONY: generate test

test:
	go test ./...
generate:
	cd src && go generate
