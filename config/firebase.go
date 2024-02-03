package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func NewFirebase(
	credentialsPath,
	databaseURL string,
) (databaseClient *db.Client, err error) {
	var (
		opt  = option.WithCredentialsFile(credentialsPath)
		conf = &firebase.Config{DatabaseURL: databaseURL}
	)

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
		return
	}

	databaseClient, err = app.Database(context.Background())
	if err != nil {
		log.Fatalf("error creating Firebase DB client: %v", err)
		return
	}

	return
}
