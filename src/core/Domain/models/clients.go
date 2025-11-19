package domain_models

import "time"

// GrantTypes  client_credentials
type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SecretKey string `json:"secret_key"`

	//
	AiCommands []string `json:"ai_commands"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
}
