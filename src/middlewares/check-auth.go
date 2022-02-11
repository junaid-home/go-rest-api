package middlewares

import (
	"net/http"
	"strings"

	"api/src/helpers"
)

var error = helpers.CustomError{}

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) < 2 {
			error.ApiError(w, http.StatusForbidden, "Token not provided!")
			return
		}

		token := bearerToken[1]

		_, err := helpers.VerifyJwtToken(token)
		if err != nil {
			error.ApiError(w, http.StatusForbidden, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
