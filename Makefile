postgres:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgresdb dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgres://root:root@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:root@localhost:5432/simplebank?sslmode=disable" -verbose up
sqlc:
	sqlc generate
.PHONY:
	postgres createdb dropdb migratedown migrateup