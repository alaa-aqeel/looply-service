package services

import domain_commands "github.com/alaa-aqeel/looply-app/src/core/Domain/commands"

type AuthService struct {
}

func (s *AuthService) NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) ChcekAuthention(cmd *domain_commands.AuthCommand) bool {

	return true
}
