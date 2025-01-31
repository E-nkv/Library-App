package api

import (
	"encoding/json"
	"fmt"
	"io"
	"library/db/types"
	"library/errs"
	"net/http"
	"strconv"
)

func (app App) Handle_GetUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	if idStr == "" {
		WriteJsonError(w, http.StatusBadRequest, "request did not specify an id for the user")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
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

	//unmarshal the body only if there is one. if theres not, then take the 0 values (meaning no limit nor lastID where passed)
	if len(b) > 0 {
		if err := json.Unmarshal(b, &reqBody); err != nil {
			WriteJsonError(w, http.StatusBadRequest, "invalid request body")
			return
		}
	}

	us, err := app.Models.Users.GetUsers(reqBody.Limit, reqBody.LastID)
	if err != nil {
		fmt.Println(err)
		WriteJsonServerError(w)
		return
	}
	WriteJsonResp(w, http.StatusOK, us, "users")
}

func (app App) Handle_CreateUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonServerError(w)
		return
	}
	var u types.UserCreate
	if err := json.Unmarshal(b, &u); err != nil {
		WriteJsonError(w, http.StatusBadRequest, "bad request")
		return
	}
	id, err := app.Models.Users.CreateUser(&u)
	if err != nil {
		switch err {
		case errs.ErrDuplicateEmail:
			WriteJsonError(w, http.StatusBadRequest, "email already exists")
		default:
			WriteJsonServerError(w)
		}
		return
	}
	WriteJsonResp(w, http.StatusOK, id, "user_id")
}

func (app App) Handle_DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		WriteJsonError(w, http.StatusBadRequest, "specify the id")
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		WriteJsonError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if err := app.Models.Users.DeleteUser(id); err != nil {
		switch err {
		case errs.ErrNotFound:
			WriteJsonError(w, http.StatusBadRequest, err.Error())
		default:
			WriteJsonServerError(w)
		}
		return
	}
	WriteJsonResp(w, http.StatusOK, nil, "")

}
