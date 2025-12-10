package client_repo

import (
	"fmt"

	domain_models "github.com/alaa-aqeel/looply-app/src/core/Domain/models"
	"github.com/jackc/pgx/v5"
)

func (r *ClientRepo) scan(rows pgx.Row) (*domain_models.Client, error) {
	var client domain_models.Client
	err := rows.Scan(
		&client.ID,
		&client.Name,
		&client.SecretKey,
		&client.AiCommands,
		&client.Active,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan client: %w", err)
	}

	return &client, nil
}

func (r *ClientRepo) scanList(rows pgx.Rows) ([]*domain_models.Client, error) {

	var listOfClients []*domain_models.Client
	for rows.Next() {
		client, err := r.scan(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan client: %w", err)
		}
		listOfClients = append(listOfClients, client)
	}

	return listOfClients, nil
}
