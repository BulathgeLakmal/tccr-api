postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpassword -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root course

dropdb:
	docker exec -it postgres12 dropdb course

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/course?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/course?sslmode=disable" -verbose up 1 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/course?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/course?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY:
	postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server