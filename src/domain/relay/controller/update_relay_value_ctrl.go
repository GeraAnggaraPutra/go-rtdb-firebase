package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"go-firebase/constant"
	"go-firebase/src/domain/relay/model"
	"go-firebase/src/module"
)

func (ctrl *RelayController) UpdateRelayValueCtrl() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var req model.UpdateRelayRequest

		if err = c.Bind(&req); err != nil {
			err = fmt.Errorf(constant.ErrFailedParseRequest.Error(), err.Error())
			return
		}

		relayNumberStr := c.Param("relay-number")

		relayNumber, err := strconv.Atoi(relayNumberStr)
		if err != nil {
			return module.ResponseError(c, http.StatusBadRequest, constant.ErrFailedParseRequest, "invalid relay number")
		}

		if relayNumber == 0 {
			return module.ResponseError(c, http.StatusBadRequest, constant.ErrFailedParseRequest, "missing relay number")
		}

		if relayNumber < 12 || relayNumber > 14 {
			return module.ResponseError(c, http.StatusBadRequest, constant.ErrFailedParseRequest, "relay number must be 12, 13 or 14")
		}

		if req.Value < 0 || req.Value > 1 {
			return module.ResponseError(c, http.StatusBadRequest, constant.ErrFailedParseRequest, "value must be between 0 and 1")
		}

		err = ctrl.service.UpdateRelayValueSvc(c.Request().Context(), relayNumber, req.Value)
		if err != nil {
			return module.ResponseError(c, http.StatusInternalServerError, err.Error(), "failed update relay value")
		}

		return module.ResponseData(c, http.StatusOK, nil, fmt.Sprintf("successfully updated relay %d with value %d", relayNumber, req.Value))
	}
}
