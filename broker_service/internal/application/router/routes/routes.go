package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/controller"
)

var ctlr controller.Controller

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
	AdminAccess    bool
}

func Start() {
	ctlr = controller.New()
}
