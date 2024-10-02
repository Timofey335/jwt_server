package user

import (
	"time"

	"github.com/Timofey335/jwt_server/internal/repository"
	def "github.com/Timofey335/jwt_server/internal/service"
)

const (
	refreshTokenSecretKey  = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	refreshTokenExpiration = 60 * time.Minute
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *serv {
	return &serv{
		userRepository: userRepository,
	}
}
