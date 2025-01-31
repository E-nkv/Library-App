package main

import (
	"fmt"
	"library/cmd/api"
	"library/db/models"
	"log"
	"net/http"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	app := api.App{Models: *models.NewModels(db)}
	router := app.Routes()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("started server on :8080")
	log.Fatal(server.ListenAndServe())
}
