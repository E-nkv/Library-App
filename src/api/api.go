package api

import (
	"database/sql"
	"fmt"
	"net/http"

	dbModels "github.com/E-nkv/libraryAPI/src/database/models"
)

type App struct {
	Models *dbModels.Models
}

func NewApp(dbConn *sql.DB) *App {
	return &App{Models: dbModels.NewModels(dbConn)}
}

func (app *App) handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthceck")
}

func (app *App) Routes() http.Handler {
	userHandler := app.genUserHandler()
	booksHandler := app.genBookHandler()
	authorsHandler := app.genAuthorHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello index")
	})
	mux.HandleFunc("/healthcheck", app.handleHealthcheck)

	mux.Handle("/users/", http.StripPrefix("/users", userHandler))
	mux.Handle("/books/", http.StripPrefix("/books", booksHandler))
	mux.Handle("/authors/", http.StripPrefix("/authors", authorsHandler))
	return mux
}

func (app *App) genUserHandler() http.Handler {
	userHandler := http.NewServeMux()
	userHandler.HandleFunc("GET /", app.handleUsers_getAll)
	return userHandler

}
func (app *App) genBookHandler() http.Handler {
	userHandler := http.NewServeMux()
	userHandler.HandleFunc("GET /", helloFrom("/books/"))
	return userHandler

}
func (app *App) genAuthorHandler() http.Handler {
	userHandler := http.NewServeMux()
	userHandler.HandleFunc("GET /", helloFrom("/authors/"))
	return userHandler
}

func helloFrom(s string) func(w http.ResponseWriter, r *http.Request) {
	s = fmt.Sprintf("Hello from %s!", s)

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(s)
		w.Write([]byte(s))
	}
}
