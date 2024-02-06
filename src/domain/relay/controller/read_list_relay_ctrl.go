package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-firebase/src/module"
)

func (ctrl *RelayController) ReadListRelayCtrl() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		data, err := ctrl.service.ReadListRelaySvc(c.Request().Context())
		if err != nil {
			return module.ResponseError(c, http.StatusInternalServerError, err.Error(), "failed get relay state list")
		}

		return module.ResponseData(c, http.StatusOK, data, "successfully read relay state list")
	}
}
