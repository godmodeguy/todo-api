.SLIENT:
.PHONY:

build:
	go build cmd/main.go

run:
	go run cmd/main.go

fmt:
	go mod tidy
	go fmt learn/todoapi/...
