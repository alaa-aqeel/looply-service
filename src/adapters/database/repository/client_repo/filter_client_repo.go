package client_repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/alaa-aqeel/looply-app/src/core/ports"
)

func (r *ClientRepo) query() squirrel.SelectBuilder {
	return database.SqlBuilder.
		Select("id, name, secret_key, active, ai_commands, created_at, updated_at").
		From("clients")
}

func (r *ClientRepo) filter(filter ports.FilterClients) squirrel.SelectBuilder {
	query := r.query()
	if filter.Name.IsSet {
		query = query.Where(squirrel.ILike{"name": "%" + filter.Name.Value + "%"})
	}
	if filter.Active.IsSet {
		query = query.Where(squirrel.Eq{"active": filter.Active.Value})
	}

	return query
}

func (r *ClientRepo) FindById(ctx context.Context, id string) (*domain_models.Client, error) {

	sql, args, err := r.query().Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRow(ctx, sql, args...)
	admin, err := r.scan(row)
	if err != nil {

		return nil, err
	}

	return admin, nil
}

func (r *ClientRepo) ClientExists(ctx context.Context, id, secretKey string) (bool, error) {

	sql, args, err := r.query().Where(squirrel.Eq{"id": id, "secret_key": secretKey}).ToSql()
	if err != nil {
		return false, err
	}

	row := r.db.QueryRow(ctx, sql, args...)
	_, err = r.scan(row)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ClientRepo) GetAll(ctx context.Context, filter ports.FilterClients) ([]*domain_models.Client, error) {
	sql, args, err := r.filter(filter).
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return r.scanList(rows)
}
