package repository

import (
	"context"

	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
)

type BaseRepo[T any] struct {
	db *database.Db
}

func NewBaseRepo[T any](db *database.Db) *BaseRepo[T] {
	return &BaseRepo[T]{
		db: db,
	}
}

func (r *BaseRepo[T]) FindById(ctx context.Context, table, cols, id string) (*T, error) {
	database.SqlBuilder.
		Select(cols).
		From(table)

	return nil, nil
}
