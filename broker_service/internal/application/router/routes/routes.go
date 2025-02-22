package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/controller"
)

var ctlr controller.Controller

type Authentication string

const (
	ADMIN   Authentication = "admin"
	DEFAULT Authentication = "default"
	NONE    Authentication = "none"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication Authentication
}

func Start() {
	ctlr = controller.New()
}
