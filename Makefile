server:
	go run cmd/main.go

test:
	go test -v -cover ./...

_PHONY: server test