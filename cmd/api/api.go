package api

import (
	"library/db/models"
	"net/http"
)

type App struct {
	Models models.Models
}

func (a App) Routes() http.Handler {
	h := http.NewServeMux()
	uh := a.genUserHandler()
	h.HandleFunc("/api/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		WriteJsonResp(w, http.StatusOK, "healthcheck", "data")
	})
	h.Handle("/api/users/", http.StripPrefix("/api/users", uh))
	return h
}

func (app App) genUserHandler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("GET /", app.Handle_GetUsers)
	h.HandleFunc("GET /{id}", app.Handle_GetUser)

	return h
}
