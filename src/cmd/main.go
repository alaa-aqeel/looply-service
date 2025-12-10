package main

import (
	"context"
	"fmt"
	"log"
	"os"

	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
	"github.com/alaa-aqeel/looply-app/src/adapters/database/repository"
	"github.com/alaa-aqeel/looply-app/src/adapters/logger"
	domain_commands "github.com/alaa-aqeel/looply-app/src/core/Domain/commands"
	"github.com/alaa-aqeel/looply-app/src/core/services"
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

	client, err := service.ClientService().CreateClient(domain_commands.CreateClientCommand{
		Name: "Test Client 7",
		AiCommands: []string{
			"hi", "no",
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client.Name)
}
