package utils

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			RespondWithJSON(w, http.StatusBadRequest, "Bad Request")
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			RespondWithJSON(w, http.StatusBadRequest, "Bad Request")
			return
		}
		if !tkn.Valid {
			RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
