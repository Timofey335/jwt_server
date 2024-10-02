package service

import (
	"context"

	descAuth "github.com/Timofey335/jwt_server/pkg/auth_v1"
)

type UserService interface {
	Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error)
	GetRefreshToken(ctx context.Context, req *descAuth.GetRefreshTokenRequest) (*descAuth.GetRefreshTokenResponse, error)
}
