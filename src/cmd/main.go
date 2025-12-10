package main

import (
	"context"
	"fmt"
	"log"
	"os"

	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
	"github.com/alaa-aqeel/looply-app/src/adapters/database/repository"
	"github.com/alaa-aqeel/looply-app/src/adapters/logger"
	"github.com/alaa-aqeel/looply-app/src/core/services"
	"github.com/alaa-aqeel/looply-app/src/shared"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("[godotenv]: Error loading .env file")
	}

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal("[logger]: " + err.Error())
	}

	db := database.NewDatabase(logger)
	err = db.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("[database]: " + err.Error())
	}
	repo := repository.NewRepoContainer(db, logger)
	service := services.NewServiceContainer(repo, logger)

	clients, err := service.ClientService().GetAll(services.ArgsClients{
		Limit:  shared.SetValue[int64](1),
		Page:   shared.SetValue[int64](1),
		Name:   shared.NilValue[string](),
		Active: shared.NilValue[bool](),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(clients)
}
