package services

import (
	"context"

	domain_commands "github.com/alaa-aqeel/looply-app/src/core/Domain/commands"
	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/alaa-aqeel/looply-app/src/core/ports"
	"github.com/google/uuid"
)

type ClientService struct {
	repo ports.ClientRepoPort
}

func NewClientService(repo ports.ClientRepoPort) *ClientService {

	return &ClientService{
		repo: repo,
	}
}

func (s *ClientService) CreateClient(cmd domain_commands.CreateClientCommand) (*domain_models.Client, error) {

	return s.repo.Create(context.Background(), &domain_models.Client{
		Name:       cmd.Name,
		AiCommands: cmd.AiCommands,
		SecretKey:  uuid.NewString(),
		Active:     true,
	})
}
