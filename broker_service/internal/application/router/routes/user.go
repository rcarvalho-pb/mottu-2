package routes

import (
	"fmt"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/middleware"
)

func ConfigUserRoutes(mux *http.ServeMux) {
	for _, r := range UserRoutes {
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

var UserRoutes = []Route{
	{
		Uri:            "/create-user",
		Method:         http.MethodPost,
		Function:       ctlr.UserController.CreateUser,
		Authentication: false,
		AdminAccess:    false,
	},
}
