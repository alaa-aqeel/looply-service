package client_repo

import (
	"context"
	"time"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
)

func (r *ClientRepo) Save(ctx context.Context, client *domain_models.Client) error {
	client.UpdatedAt = time.Now()
	err := r.base.Update(ctx, "clients", client.ID, map[string]any{
		"name":        client.Name,
		"active":      client.Active,
		"ai_commands": client.AiCommands,
		"updated_at":  client.UpdatedAt,
	})

	return err
}

func (r *ClientRepo) UpdateSecretKey(ctx context.Context, id string, secretKey string) error {
	err := r.base.Update(ctx, "clients", id, map[string]any{
		"secret_key": secretKey,
		"updated_at": time.Now(),
	})

	return err
}

func (r *ClientRepo) UpdateActive(ctx context.Context, id string, active bool) error {
	err := r.base.Update(ctx, "clients", id, map[string]any{
		"active":     active,
		"updated_at": time.Now(),
	})

	return err
}
