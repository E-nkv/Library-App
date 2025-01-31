package api

import (
	"encoding/json"
	"io"
	"library/db/types"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app App) Handle_LoginWithCreds(w http.ResponseWriter, r *http.Request) {
	//do not allow already logged in users to log in
	if cookie, err := r.Cookie("AuthCookie"); err == nil {
		tokenStr := cookie.Value
		_, err := parseJWT(tokenStr)
		if err == nil {
			WriteJsonError(w, http.StatusBadRequest, "cannot login an already logged in user")
			return
		}
	}

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonServerError(w)
		return
	}
	var u types.UserLogin
	if err := json.Unmarshal(bs, &u); err != nil {
		WriteJsonError(w, http.StatusBadRequest, "malformed json data")
		return
	}
	userDb, err := app.Models.Users.GetUserByEmail(u.Email)
	if err != nil {
		WriteJsonError(w, http.StatusBadRequest, "User does not exist. Please sign up first") //
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDb.HashPass), []byte(u.PlainPasswd)); err != nil {
		WriteJsonError(w, http.StatusBadRequest, "invalid email or password")
		return
	}
	tokenStr, err := genJWT(userDb.ID, userDb.Role)
	if err != nil {

		WriteJsonServerError(w)
		return
	}

	cookie := http.Cookie{
		Name:     "AuthCookie",
		Value:    tokenStr,
		HttpOnly: false,
		Secure:   false,
	}
	http.SetCookie(w, &cookie)

	WriteJsonResp(w, http.StatusOK, "logged in user successfully", "success")

}

func (app App) Handle_Logout(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.Context().Value("userAuth").(types.UserAuth); !ok {
		//technically shouldnt happen, since the middleware already put the userAuth in the r.ctx
		WriteJsonServerError(w)
		return
	}
	cookie := http.Cookie{
		Name:   "AuthCookie",
		Value:  "",
		MaxAge: -1, //delete the cookie
	}
	http.SetCookie(w, &cookie)
	WriteJsonResp(w, http.StatusAccepted, "success", "logoutStatus")
}
