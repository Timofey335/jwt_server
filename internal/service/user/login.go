package user

import (
	"context"
	"errors"

	"github.com/Timofey335/jwt_server/internal/repository/model"
	"github.com/Timofey335/jwt_server/internal/utils"
	descAuth "github.com/Timofey335/jwt_server/pkg/auth_v1"
)

func (s *serv) Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	user, err := s.userRepository.GetUserData(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	//todo: to add hashing password

	if req.Password != user.Password {
		return nil, errors.New("password incorrect")
	}

	refreshToken, err := utils.GenerateToken(model.UserData{
		Username: req.Username,
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
