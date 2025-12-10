package client_service_test

import (
	"context"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/alaa-aqeel/looply-app/src/core/ports"
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

func (m *MockClientRepo) FindById(ctx context.Context, id string) (*domain_models.Client, error) {
	return nil, nil
}

func (m *MockClientRepo) ClientExists(ctx context.Context, id, secretKey string) (bool, error) {
	return false, nil
}

func (m *MockClientRepo) GetAll(ctx context.Context, filter ports.FilterClients) ([]*domain_models.Client, error) {
	return nil, nil
}

func (m *MockClientRepo) Update(ctx context.Context, client *domain_models.Client) error {
	return nil
}

func (m *MockClientRepo) UpdateActive(ctx context.Context, id string, active bool) error {
	return nil
}

func (m *MockClientRepo) UpdateSecretKey(ctx context.Context, id string, secretKey string) error {
	return nil
}

func (m *MockClientRepo) Delete(ctx context.Context, id string) error {
	return nil
}
