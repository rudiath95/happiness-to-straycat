include .$(PWD)/.env

say_hello:
    echo "TEST"

postgres:
	docker run --name some-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it some-postgres createdb --username=root --owner=root happiness-to-straycat

dropdb:
	docker exec -it some-postgres dropdb happiness-to-straycat

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/happiness-to-straycat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/happiness-to-straycat?sslmode=disable" -verbose down

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc

