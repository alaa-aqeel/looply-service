package base_repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
)

type Repo[T any] struct {
	Db *database.Db
}

func (r *Repo[T]) Insert(ctx context.Context, table string, data map[string]any) error {
	sql, args, err := database.SqlBuilder.
		Insert(table).
		SetMap(data).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return err
	}

	err = r.Db.Exec(ctx, sql, args...)
	if err != nil {
		return database.MapPgError(err)
	}

	return nil
}

func (r *Repo[T]) Update(ctx context.Context, table string, id string, data map[string]any) error {
	sql, args, err := database.SqlBuilder.
		Update(table).
		SetMap(data).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	err = r.Db.Exec(ctx, sql, args...)
	if err != nil {
		return database.MapPgError(err)
	}

	return nil
}

func (r *Repo[T]) Delete(ctx context.Context, table string, id string) error {
	sql, args, err := database.SqlBuilder.
		Delete(table).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	err = r.Db.Exec(ctx, sql, args...)
	if err != nil {
		return database.MapPgError(err)
	}

	return nil
}
