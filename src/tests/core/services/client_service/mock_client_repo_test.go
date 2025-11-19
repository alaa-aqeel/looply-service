package client_service_test

import (
	"context"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
)

type MockClientRepo struct {
	// You can define behavior for each method
	CreateFunc func(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error)
}

func (m *MockClientRepo) Create(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, client)
	}
	// default behavior
	return client, nil
}
