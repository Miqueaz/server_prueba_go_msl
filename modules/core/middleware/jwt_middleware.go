package middleware

import (
	token "main/security/token"
	"net/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		t := req.Header.Get("Autorization")
		_, err := token.ValidToken(t)
		if err != nil {
			sendJSON(res, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(res, req)
	})
}
