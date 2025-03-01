package middleware

import (
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func JWTMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		// If empty value
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			}

			helper.WriteToResponseBody(w, webResponse)
			return
		}

		// if invalid token
		tokenString := authHeader[len("Bearer "):]
		err := VerifyToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			}

			helper.WriteToResponseBody(w, webResponse)
			return
		}

		next(w, r, ps)
	}
}
