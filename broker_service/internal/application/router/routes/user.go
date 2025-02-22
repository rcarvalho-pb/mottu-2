package routes

import (
	"fmt"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/middleware"
)

func ConfigUserRoutes(mux *http.ServeMux) {
	for _, r := range UserRoutes {
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

var UserRoutes = []Route{
	{
		Uri:            "/create-user",
		Method:         http.MethodPost,
		Function:       ctlr.UserController.CreateUser,
		Authentication: NONE,
	},
	{
		Uri:            "/",
		Method:         http.MethodGet,
		Function:       ctlr.UserController.GetAllActiveUsers,
		Authentication: ADMIN,
	},
	{
		Uri:            "/all",
		Method:         http.MethodGet,
		Function:       ctlr.UserController.GetAllUsers,
		Authentication: DEFAULT,
	},
	{
		Uri:            "/{userId}",
		Method:         http.MethodGet,
		Function:       ctlr.UserController.GetUserById,
		Authentication: ADMIN,
	},
	{
		Uri:            "/{userId}/update",
		Method:         http.MethodPut,
		Function:       ctlr.UserController.UpdateUser,
		Authentication: NONE,
	},
	{
		Uri:            "/{userId}/update-password",
		Method:         http.MethodPatch,
		Function:       ctlr.UserController.UpdatePassword,
		Authentication: NONE,
	},
	{
		Uri:            "/{userId}/deactivate-user",
		Method:         http.MethodPatch,
		Function:       ctlr.UserController.DeactivateUser,
		Authentication: NONE,
	},
	{
		Uri:            "/{userId}/reactivate-user",
		Method:         http.MethodPatch,
		Function:       ctlr.UserController.ReactivateUser,
		Authentication: NONE,
	},
}
