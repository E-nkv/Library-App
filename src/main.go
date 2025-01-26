package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/E-nkv/libraryAPI/src/api"
)

func main() {
	var dbConn *sql.DB = nil

	app := api.NewApp(dbConn)
	router := app.Routes()
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Printf("Started server at %s\n", "localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
