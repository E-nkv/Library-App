package api

import (
	"library/db/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Models models.Models
}

func (app App) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		WriteJsonResp(w, http.StatusOK, "hello api!", "hello")
	})

	r.Get("/api/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		WriteJsonResp(w, http.StatusOK, "healthcheck", "data")
	})
	r.Route("/api/users", app.usersRouter)
	r.Post("/api/login", app.Handle_LoginWithCreds)
	r.With(AuthOnlyMdw).Get("/api/logout", app.Handle_Logout)

	return r
}

func (app App) usersRouter(r chi.Router) {
	r.Use(AuthOnlyMdw)
	r.Get("/", app.Handle_GetUsers)
	r.Get("/{id}", app.Handle_GetUser)
	r.Post("/", app.Handle_CreateUser)
	r.Delete("/{id}", app.Handle_DeleteUser)
}
