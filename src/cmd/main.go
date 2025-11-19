package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alaa-aqeel/looply-app/src/adapters/database/logger"
	database "github.com/alaa-aqeel/looply-app/src/adapters/database/pgsql"
	"github.com/alaa-aqeel/looply-app/src/adapters/database/repository"
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
		Name: "test client 2",
		AiCommands: []string{
			"do that me",
			"new how are ur",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client.ID)
	fmt.Println(client.SecretKey)
	fmt.Println(client.Name)

	// d132a384-9af9-4bf0-8c15-423f4e39e32c
	// 0fa5f2f3-7ade-452a-a808-7dbe51717405
	// test client 2
}
