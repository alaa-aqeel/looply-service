package client_repo

import (
	"context"
	"time"

	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/google/uuid"
)

func (r *ClientRepo) Create(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error) {
	client.ID = uuid.NewString()
	client.CreatedAt = time.Now()
	client.UpdatedAt = client.CreatedAt

	sql, args, err := database.SqlBuilder.
		Insert("clients").
		Columns("id, name, secret_key, active, ai_commands, created_at, updated_at").
		Values(client.ID, client.Name, client.SecretKey, client.Active, client.AiCommands, client.CreatedAt, client.UpdatedAt).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, sql, args...).Scan(&client.ID)
	if err != nil {
		return nil, database.MapPgError(err)
	}

	return client, nil
}
