package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	emddw "github.com/labstack/echo/v4/middleware"

	relayRoute "go-firebase/src/domain/relay/routes"
	"go-firebase/src/middleware"
	"go-firebase/src/module"
)

func Routes(e *echo.Echo, m *module.Module) {
	e.HideBanner = true

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		log.Printf("Error: %v", err)
		e.DefaultHTTPErrorHandler(err, ctx)
	}

	e.Use(emddw.LoggerWithConfig(emddw.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Response().Status < http.StatusBadRequest
		},
	}))

	// handle middleware
	middleware.CorsMiddleware(e)
	middleware.RecoverMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": fmt.Sprintf("%s API is Running", os.Getenv("APP_NAME")),
		})
	})

	e.Static("/media", "src/public")

	relayRoute.AddRelayRoute(e, m)
}
