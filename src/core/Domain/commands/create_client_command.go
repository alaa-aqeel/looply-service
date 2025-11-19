package domain_commands

type CreateClientCommand struct {
	Name       string   `json:"name" form:"name"`
	AiCommands []string `json:"ai_commands" form:"ai_commands"`
}
