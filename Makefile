.PHONY:

build:
	go build cmd/main.go

run:
	go run cmd/main.go

fmt:
	go mod tidy
	go fmt learn/todoapi/...

db:
	docker run \
		--name todo-db \
		-e POSTGRES_USER=usr \
		-e POSTGRES_PASSWORD=pwd \
		-p 5432:5432 \
		-d --rm \
		-v tododb:/var/lib/postgresql/data \
		postgres

migrate-create:
	# rm -ri schema
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://usr:pwd@localhost:5432?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://usr:pwd@localhost:5432?sslmode=disable' down

migrate-force-first:
	migrate -path ./schema -database 'postgres://usr:pwd@localhost:5432?sslmode=disable' force 000001

db-shell:
	docker exec -it todo-db /bin/bash
