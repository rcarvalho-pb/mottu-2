package routes

import (
	"fmt"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/middleware"
)

func ConfigAuthRoutes(mux *http.ServeMux) {
	for _, r := range AuthRoutes {
		handler := middleware.Logger(r.Function)
		switch r.Authentication {
		case ADMIN:
			handler = middleware.Logger(middleware.IsAdmin(r.Function))
		case DEFAULT:
			handler = middleware.Logger(middleware.Authenticate(r.Function))
		}
		mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Uri), handler)
	}
}

var AuthRoutes = []Route{
	{
		Uri:            "/",
		Method:         http.MethodPost,
		Function:       ctlr.AuthController.Authenticate,
		Authentication: NONE,
	},
}
