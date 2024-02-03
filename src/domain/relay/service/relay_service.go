package service

import "firebase.google.com/go/v4/db"

type RelayService struct {
	dbClient *db.Client
}

func NewRelayService(dbClient *db.Client) *RelayService {
	return &RelayService{
		dbClient: dbClient,
	}
}
