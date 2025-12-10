package client_repo

import (
	"context"
	"time"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/google/uuid"
)

func (r *ClientRepo) Create(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error) {
	client.ID = uuid.NewString()
	client.CreatedAt = time.Now()
	client.UpdatedAt = client.CreatedAt
	err := r.base.Insert(ctx, "clients", map[string]any{
		"id":          client.ID,
		"name":        client.Name,
		"secret_key":  client.SecretKey,
		"active":      client.Active,
		"ai_commands": client.AiCommands,
		"created_at":  client.CreatedAt,
		"updated_at":  client.UpdatedAt,
	})

	return client, err
}
