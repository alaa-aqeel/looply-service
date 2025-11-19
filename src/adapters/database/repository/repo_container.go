package repository

import (
	"sync"

	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
	"github.com/alaa-aqeel/looply-app/src/adapters/database/repository/client_repo"
	"github.com/alaa-aqeel/looply-app/src/core/ports"
)

type RepoContainer struct {
	db     *database.Db
	onces  map[string]*sync.Once
	mu     sync.Mutex
	logger ports.LoggerPort

	// Repositories
	clientRepo *client_repo.ClientRepo
}

func NewRepoContainer(db *database.Db, logger ports.LoggerPort) *RepoContainer {
	return &RepoContainer{
		db:     db,
		logger: logger,
		onces:  make(map[string]*sync.Once),
	}
}

func (c *RepoContainer) once(name string) *sync.Once {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.onces[name]; !ok {
		c.onces[name] = &sync.Once{}
	}

	return c.onces[name]
}

func (c *RepoContainer) ClientRepo() *client_repo.ClientRepo {
	c.once("admin_repo").Do(func() {
		c.clientRepo = client_repo.NewClientRepo(c.db)
	})

	return c.clientRepo
}
