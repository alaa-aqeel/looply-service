package database

import (
	"context"
	"sync"

	"github.com/alaa-aqeel/looply-app/src/core/ports"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	instance *Db
	once     sync.Once
)

type Db struct {
	pool   *pgxpool.Pool
	logger ports.LoggerPort
}

func NewDatabase(logger ports.LoggerPort) *Db {
	once.Do(func() {
		instance = &Db{
			logger: logger,
		}
	})
	return instance
}

func (p *Db) Connect(ctx context.Context, dbUrl string) error {
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		p.logger.Log(ports.Error, ports.Logger{
			Message: "connect to database failed",
			Tag:     "database:connect:failed",
			Error:   err,
			Args: map[string]string{
				"db_url": dbUrl,
			},
		})
		return MapPgError(err)
	}
	p.pool = pool
	p.logger.Log(ports.Info, ports.Logger{
		Message: "connect to database success",
		Tag:     "database:connect:success",
		Error:   err,
		Args: map[string]string{
			"db_url": dbUrl,
		},
	})

	return nil
}

func (p *Db) Close() {

	p.pool.Close()
}

func (p *Db) Pool() *pgxpool.Pool {

	return p.pool
}

func (p *Db) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {

	return p.Pool().QueryRow(ctx, sql, args...)
}

func (p *Db) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	row, err := p.Pool().Query(ctx, sql, args...)
	return row, MapPgError(err)
}

// Exec executes a statement with arguments
func (p *Db) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := p.Pool().Exec(ctx, sql, args...)

	return MapPgError(err)
}
