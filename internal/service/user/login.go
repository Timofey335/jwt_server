package user

import (
	"context"
	"errors"
	"time"

	"github.com/Timofey335/jwt_server/internal/repository/model"
	"github.com/Timofey335/jwt_server/internal/utils"
	descAuth "github.com/Timofey335/jwt_server/pkg/auth_v1"
)

const (
	refreshTokenSecretKey  = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	refreshTokenExpiration = 60 * time.Minute
)

func (s *serv) Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	user, err := s.userRepository.GetPassword(ctx, req)
	if err != nil {
		return nil, err
	}

	//todo: to add hashing password

	if req.Password != user.Password {
		return nil, errors.New("password incorrect")
	}

	refreshToken, err := utils.GenerateToken(model.UserData{
		Username: user.Name,
		Role:     user.Role,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &descAuth.LoginResponse{
		RefreshToken: refreshToken,
	}, nil
}
