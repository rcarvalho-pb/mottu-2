package router

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/router/routes"
)

func New() *http.ServeMux {
	routes.Start()
	router := http.NewServeMux()

	router.Handle("/api/user/", http.StripPrefix("/api/user", newUserMux()))
	router.Handle("/api/auth/", http.StripPrefix("/api/auth", newAuthMux()))

	return router
}

func newUserMux() *http.ServeMux {
	router := http.ServeMux{}
	routes.ConfigUserRoutes(&router)
	return &router
}

func newAuthMux() *http.ServeMux {
	router := http.ServeMux{}
	routes.ConfigAuthRoutes(&router)
	return &router
}
