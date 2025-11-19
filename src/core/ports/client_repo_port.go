package ports

import (
	"context"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
)

type ClientRepoPort interface {
	Create(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error)
}
