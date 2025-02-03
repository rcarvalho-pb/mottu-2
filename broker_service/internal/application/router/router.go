package router

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/router/routes"
)

func New() *http.ServeMux {
	routes.Start()
	router := http.NewServeMux()
	router.Handle("/api/user/", http.StripPrefix("/api/user", newUserMux()))
	return router
}

func newUserMux() *http.ServeMux {
	mux := http.NewServeMux()
	routes.ConfigUserRoutes(mux)
	return mux
}
