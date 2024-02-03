package module

import (
	"firebase.google.com/go/v4/db"
)

type Module struct {
	db *db.Client
}

func NewModule(
	database *db.Client,
) *Module {
	return &Module{
		db: database,
	}
}

func (s *Module) GetDBFirebase() *db.Client {
	return s.db
}
