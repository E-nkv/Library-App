package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dbConnStr := "postgres://postgres:admin@localhost:5432/libraryapp?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		panic(err)
	}
	seeder := S{DB: db}
	if err := seeder.seedDb(); err != nil {
		panic(err)
	}

	r := seeder.DB.QueryRow("SELECT email, hash_pass FROM USERS WHERE id = $1", 10)

	var e, p string

	err = r.Scan(&e, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(e, " ", p)

}

type S struct {
	DB *sql.DB
}

func (s S) seedDb() error {
	fmt.Println("Starting db seeding")
	err := s.seedUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println("Db seeding ended succesfuly")
	return nil
}
