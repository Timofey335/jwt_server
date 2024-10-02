package repository

import (
	"context"

	repoModel "github.com/Timofey335/jwt_server/internal/repository/model"
)

type UserRepository interface {
	GetUserData(ctx context.Context, userName string) (*repoModel.UserRepoModel, error)
}
