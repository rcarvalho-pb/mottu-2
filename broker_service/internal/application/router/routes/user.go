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
	{
		Uri:            "/",
		Method:         http.MethodGet,
		Function:       ctlr.UserController.GetAllActiveUsers,
		Authentication: false,
		AdminAccess:    false,
	},
	{
		Uri:            "/all",
		Method:         http.MethodGet,
		Function:       ctlr.UserController.GetAllUsers,
		Authentication: false,
		AdminAccess:    false,
	},
	{
		Uri:            "/{userId}",
		Method:         http.MethodGet,
		Function:       ctlr.UserController.GetUserById,
		Authentication: false,
		AdminAccess:    false,
	},
	{
		Uri:            "/{userId}/update",
		Method:         http.MethodPut,
		Function:       ctlr.UserController.UpdateUser,
		Authentication: false,
		AdminAccess:    false,
	},
	{
		Uri:            "/{userId}/update-password",
		Method:         http.MethodPatch,
		Function:       ctlr.UserController.UpdatePassword,
		Authentication: false,
		AdminAccess:    false,
	},
	{
		Uri:            "/{userId}/deactivate-user",
		Method:         http.MethodPatch,
		Function:       ctlr.UserController.DeactivateUser,
		Authentication: false,
		AdminAccess:    false,
	},
	{
		Uri:            "/{userId}/reactivate-user",
		Method:         http.MethodPatch,
		Function:       ctlr.UserController.ReactivateUser,
		Authentication: false,
		AdminAccess:    false,
	},
}
