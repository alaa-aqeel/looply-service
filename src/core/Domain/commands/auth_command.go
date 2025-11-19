package domain_commands

type AuthCommand struct {
	ClientId     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
}
