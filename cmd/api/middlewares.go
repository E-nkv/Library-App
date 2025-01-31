package api

import (
	"context"
	"net/http"
)

func AuthOnlyMdw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("AuthCookie")
		if err != nil {
			WriteJsonError(w, http.StatusUnauthorized, "no auth cookie found")
			return
		}

		tokenStr := cookie.Value
		userAuth, err := parseJWT(tokenStr)
		if err != nil {

			WriteJsonError(w, http.StatusUnauthorized, "invalid or expired token. Please log in")
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "userAuth", *userAuth))
		next.ServeHTTP(w, r)
	})
}
