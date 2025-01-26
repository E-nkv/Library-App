package api

import (
	"fmt"
	"net/http"
)

func (app *App) handleUsers_getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting all users")
}
