package client_repo

import (
	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
	"github.com/alaa-aqeel/looply-app/src/adapters/database/repository/base_repo"
	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
)

type ClientRepo struct {
	db   *database.Db
	base *base_repo.Repo[domain_models.Client]
}

func NewClientRepo(db *database.Db) *ClientRepo {

	return &ClientRepo{
		db: db,
		base: &base_repo.Repo[domain_models.Client]{
			Db: db,
		},
	}
}
