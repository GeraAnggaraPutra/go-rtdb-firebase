package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"go-firebase/config"
	"go-firebase/src/api"
	"go-firebase/src/module"
)

func main() {
	// load .env file
	err := loadEnv()
	if err != nil {
		return
	}

	var (
		appPort         int
		appName         = os.Getenv("APP_NAME")
		appHost         = os.Getenv("APP_HOST")
		credentialsPath = os.Getenv("FIREBASE_CREDENTIALS_PATH")
		databaseURL     = os.Getenv("DATABASE_URL")
	)

	if port, err := strconv.Atoi(os.Getenv("APP_PORT")); port > 0 && err == nil {
		appPort = port
	}

	dbFirebase, err := config.NewFirebase(credentialsPath, databaseURL)
	if err != nil {
		return
	}

	mdl := module.NewModule(dbFirebase)

	e := echo.New()

	api.Routes(e, mdl)

	log.Printf("serving REST HTTP server || config: name=%s, host=%s, port=%d", appName, appHost, appPort)

	if err := e.Start(fmt.Sprintf("%s:%d", appHost, appPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err, "starting http server")
	}
}

func loadEnv() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error load env file : %s", err.Error())
		return
	}

	return
}
