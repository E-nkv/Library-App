m-up:
	migrate -path db/migrations -database "postgres://postgres:admin@localhost:5432/libraryapp?sslmode=disable" up

m-down:
	migrate -path db/migrations -database "postgres://postgres:admin@localhost:5432/libraryapp?sslmode=disable" down

m-ff:
	migrate -path db/migrations -database "postgres://postgres:admin@localhost:5432/libraryapp?sslmode=disable" force 2

seed:
	go run ./db/seeds

run:
	go run ./cmd