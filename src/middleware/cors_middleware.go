package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-firebase/constant"
)

func CorsMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderAccept, echo.HeaderContentType, echo.HeaderAuthorization, constant.DefaultAllowHeaderAccessToken, constant.DefaultAllowHeaderRefreshToken},
	}))
}
