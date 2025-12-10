package ports

import (
	"context"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
)

type FilterClients struct {
	Limit  uint64 `json:"limit" form:"limit"`
	Offset uint64 `json:"offset" form:"offset"`
	Name   string `json:"name" form:"name"`
	Active *bool  `json:"active" form:"active"`
}

type ClientRepoPort interface {
	FindById(ctx context.Context, id string) (*domain_models.Client, error)
	ClientExists(ctx context.Context, id, secretKey string) (bool, error)
	GetAll(ctx context.Context, filter FilterClients) ([]*domain_models.Client, error)
	Create(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error)
}
