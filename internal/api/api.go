package api

import (
	"test1/internal/database"
	"test1/models"
	"test1/storages"
)

type (
	API struct {
		Database *database.Database
		Storage  *storages.Storage[models.Object]
	}
)

func NewAPI(d *database.Database, s *storages.Storage[models.Object]) (a *API) {
	a = &API{
		Database: d,
		Storage:  s,
	}

	return
}
