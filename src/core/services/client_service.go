package services

import (
	"context"

	domain_commands "github.com/alaa-aqeel/looply-app/src/core/Domain/commands"
	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/alaa-aqeel/looply-app/src/core/ports"
	"github.com/alaa-aqeel/looply-app/src/shared"
	"github.com/google/uuid"
)

type ArgsClients struct {
	Limit  shared.Optional[int64]  `json:"limit" form:"limit"`
	Page   shared.Optional[int64]  `json:"page" form:"page"`
	Name   shared.Optional[string] `json:"name" form:"name"`
	Active shared.Optional[bool]   `json:"active" form:"active"`
}

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

func (s *ClientService) FindById(id string) (*domain_models.Client, error) {

	return s.repo.FindById(context.Background(), id)
}

func (s *ClientService) GetAll(args ArgsClients) ([]*domain_models.Client, error) {

	limit := args.Limit.ValueOrDefault(10)
	page := (args.Page.ValueOrDefault(1) - 1) * limit

	return s.repo.GetAll(context.Background(), ports.FilterClients{
		Limit:  limit,
		Offset: page,
		Name:   args.Name,
		Active: args.Active,
	})
}
