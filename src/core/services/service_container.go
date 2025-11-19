package services

import (
	"sync"

	"github.com/alaa-aqeel/looply-app/src/adapters/database/repository"
	"github.com/alaa-aqeel/looply-app/src/core/ports"
)

type ServiceContainer struct {
	onces  map[string]*sync.Once
	mu     sync.Mutex
	logger ports.LoggerPort
	repo   *repository.RepoContainer

	// Repositories
	clientService *ClientService
}

func NewServiceContainer(repo *repository.RepoContainer, logger ports.LoggerPort) *ServiceContainer {
	return &ServiceContainer{
		repo:   repo,
		logger: logger,
		onces:  make(map[string]*sync.Once),
	}
}

func (c *ServiceContainer) once(name string) *sync.Once {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.onces[name]; !ok {
		c.onces[name] = &sync.Once{}
	}

	return c.onces[name]
}

func (c *ServiceContainer) ClientService() *ClientService {
	c.once("admin_service").Do(func() {
		c.clientService = NewClientService(
			c.repo.ClientRepo(),
		)
	})

	return c.clientService
}
