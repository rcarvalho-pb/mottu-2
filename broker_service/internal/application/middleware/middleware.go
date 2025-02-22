package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/global"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/helper"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"
)

var tokenService *service.TokenService

func Init(tService *service.TokenService) {
	tokenService = tService
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := tokenService.ValidateToken(r)
		if err != nil {
			helper.ErrorJson(w, err, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), global.CLAIMS, claims)

		next(w, r.WithContext(ctx))
	}
}

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := tokenService.ValidateToken(r)
		if err != nil || claims.UserRole != "admin" {
			log.Println("erro:", err)
			helper.ErrorJson(w, fmt.Errorf("unauthorized"), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), global.CLAIMS, claims)
		next(w, r.WithContext(ctx))
	}
}
