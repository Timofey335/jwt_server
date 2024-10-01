package repository

import (
	"context"

	repoModel "github.com/Timofey335/jwt_server/internal/repository/model"
	descAuth "github.com/Timofey335/jwt_server/pkg/auth_v1"
)

type UserRepository interface {
	GetPassword(ctx context.Context, req *descAuth.LoginRequest) (*repoModel.UserRepoModel, error)
}
