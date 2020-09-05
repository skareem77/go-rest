package util

import (
	"context"
	"fmt"
	"net/http"
	h "rest-api/src/app/handler"
	"rest-api/src/app/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//JwtAuthentication validate token
var JwtAuthentication = func(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/login"}
		requestPath := r.URL.Path //current request path
		fmt.Println("jwt Authenticator")
		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				handler.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		fmt.Printf("Token Header: %v", tokenHeader)
		if tokenHeader == "" {
			h.RespondError(w, http.StatusUnauthorized, "TokenHeader: Token Authentication Failed")
			return
		}
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			h.RespondError(w, http.StatusForbidden, "Token Length:Token Authentication Failed")
			return
		}
		tokenPart := splitted[1]
		tk := &model.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("SECRET"), nil
		})

		if err != nil {
			h.RespondError(w, http.StatusForbidden, err.Error())
			return
		}

		if !token.Valid {
			h.RespondError(w, http.StatusForbidden, "Validation: Token is not valid.")
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}
