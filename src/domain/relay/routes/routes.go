package routes

import (
	"github.com/labstack/echo/v4"

	"go-firebase/src/domain/relay/controller"
	"go-firebase/src/domain/relay/service"
	"go-firebase/src/module"
)

func AddRelayRoute(e *echo.Echo, m *module.Module) {
	svc := service.NewRelayService(m.GetDBFirebase())
	ctrl := controller.NewRelayController(svc)

	rRoute := e.Group("/board1/outputs/digital")

	rRoute.GET("", ctrl.ReadListRelayCtrl())
	rRoute.PATCH("/:relay-number", ctrl.UpdateRelayValueCtrl())
}
