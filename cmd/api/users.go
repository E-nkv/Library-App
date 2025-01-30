package api

import (
	"encoding/json"
	"fmt"
	"io"
	"library/errs"
	"net/http"
	"strconv"
)

func (app App) Handle_GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got exec get user")
	idStr := r.PathValue("id")
	if idStr == "" {
		WriteJsonError(w, http.StatusBadRequest, "request did not specify an id for the user")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		WriteJsonError(w, http.StatusBadRequest, "request specified an invalid id for the user")
		return
	}

	u, err := app.Models.Users.GetUser(id)
	if err != nil {
		switch err {
		case errs.ErrNotFound:
			WriteJsonError(w, http.StatusNotFound, "the specified user was not found")
		default:
			WriteJsonServerError(w)
		}
		return
	}
	//give only the send-able data to the frontend
	output := struct {
		ID         int64  `json:"id"`
		FullName   string `json:"full_name"`
		Email      string `json:"email"`
		IsActive   bool   `json:"is_active"`
		IsVerified bool   `json:"is_verified"`
	}{
		ID: u.ID, FullName: u.FullName, Email: u.Email, IsActive: u.IsActive, IsVerified: u.IsVerified,
	}
	WriteJsonResp(w, http.StatusOK, output, "user")

}

func (app App) Handle_GetUsers(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonServerError(w)
	}
	var reqBody = struct {
		Limit  int   `json:"limit"`
		LastID int64 `json:"lastID"`
	}{}

	fmt.Println("len of b is ", len(b))
	//unmarshal the body only if there is one. if theres not, then take the 0 values (meaning no limit nor lastID where passed)
	if len(b) > 0 {
		if err := json.Unmarshal(b, &reqBody); err != nil {
			WriteJsonError(w, http.StatusBadRequest, "invalid request body")
			return
		}
	}
	fmt.Println(reqBody)
	us, err := app.Models.Users.GetUsers(reqBody.Limit, reqBody.LastID)
	if err != nil {
		WriteJsonServerError(w)
		return
	}
	WriteJsonResp(w, http.StatusOK, us, "users")
}
