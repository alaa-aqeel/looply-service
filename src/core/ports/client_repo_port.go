package ports

import (
	"context"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/alaa-aqeel/looply-app/src/shared"
)

type FilterClients struct {
	Limit  int64                   `json:"limit" form:"limit"`
	Offset int64                   `json:"offset" form:"offset"`
	Name   shared.Optional[string] `json:"name" form:"name"`
	Active shared.Optional[bool]   `json:"active" form:"active"`
}

type ClientRepoPort interface {
	FindById(ctx context.Context, id string) (*domain_models.Client, error)
	ClientExists(ctx context.Context, id, secretKey string) (bool, error)
	GetAll(ctx context.Context, filter FilterClients) ([]*domain_models.Client, error)
	Save(ctx context.Context, client *domain_models.Client) error
	UpdateActive(ctx context.Context, id string, active bool) error
	UpdateSecretKey(ctx context.Context, id string, secretKey string) error
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error)
}
