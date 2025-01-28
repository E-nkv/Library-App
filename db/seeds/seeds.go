package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dbConnStr := "postgresql://postgres:admin@localhost:5432/libraryapp?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		panic(err)
	}

	panicIfErr(err)

	seeder := S{DB: db}
	if err := seeder.seedDb(); err != nil {
		panic(err)
	}

}

type S struct {
	DB *sql.DB
}

func panicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

func (s S) seedDb() error {
	fmt.Println("Starting db seeding")

	panicIfErr(s.seedTags(10))
	panicIfErr(s.seedUsers(20))

	panicIfErr(s.seedBooks(45))

	fmt.Println("Db seeding ended succesfuly")
	return nil
}
