package client_repo

import (
	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
)

type ClientRepo struct {
	db *database.Db
}

func NewClientRepo(db *database.Db) *ClientRepo {

	return &ClientRepo{
		db: db,
	}
}
