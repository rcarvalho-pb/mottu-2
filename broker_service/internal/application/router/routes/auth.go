package routes

import (
	"fmt"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/middleware"
)

func ConfigAuthRoutes(mux *http.ServeMux) {
	for _, r := range AuthRoutes {
		if r.AdminAccess {
			mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Uri), middleware.Logger(middleware.IsAdmin(r.Function)))
		} else {
			if r.Authentication {
				mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Uri), middleware.Logger(middleware.Authenticate(r.Function)))
			} else {
				mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Uri), middleware.Logger(r.Function))
			}
		}
	}
}

var AuthRoutes = []Route{
	{
		Uri:            "/",
		Method:         http.MethodPost,
		Function:       ctlr.AuthController.Authenticate,
		Authentication: false,
		AdminAccess:    false,
	},
}
