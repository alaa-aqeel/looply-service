package client_service_test

import (
	"context"
	"errors"
	"testing"

	domain_commands "github.com/alaa-aqeel/looply-app/src/core/Domain/commands"
	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/alaa-aqeel/looply-app/src/core/services"
	"github.com/google/uuid"
)

func TestCreateClient_Success(t *testing.T) {
	id := uuid.NewString()
	mockRepo := &MockClientRepo{
		CreateFunc: func(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error) {
			client.ID = id
			return client, nil
		},
	}

	svc := services.NewClientService(mockRepo)

	created, err := svc.CreateClient(domain_commands.CreateClientCommand{
		Name:       "hello world",
		AiCommands: []string{"hello", "world"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if created.ID != id {
		t.Errorf("expected %s, got %s", id, created.ID)
	}
}

func TestCrateClient_Error(t *testing.T) {
	mockRepo := &MockClientRepo{
		CreateFunc: func(ctx context.Context, client *domain_models.Client) (*domain_models.Client, error) {

			return nil, errors.New("db error")
		},
	}

	svc := services.NewClientService(mockRepo)
	_, err := svc.CreateClient(domain_commands.CreateClientCommand{
		Name:       "hello world",
		AiCommands: []string{"hello", "world"},
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
