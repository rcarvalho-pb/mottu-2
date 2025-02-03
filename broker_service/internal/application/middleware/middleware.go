package middleware

import (
	"log"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/authentication"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/helper"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			helper.ErrorJson(w, err, http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			helper.ErrorJson(w, err, http.StatusUnauthorized)
			return
		}
		if claims, err := authentication.GetClaims(r); err != nil || claims.UserRole != "Admin" {
			helper.ErrorJson(w, err, http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
