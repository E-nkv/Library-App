package main

import (
	"database/sql"
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
	log.Fatal(srv.ListenAndServe())
}
