postgres:
	docker run --name postgres-alpine -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=12345 -d postgres:alpine
createdb:
	docker exec -it postgres-alpine createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres-alpine dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:12345@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:12345@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
./PHONY: postgres createdb dropdb migrateup migratedown sqlc