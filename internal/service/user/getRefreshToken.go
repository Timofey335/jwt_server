package user

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Timofey335/jwt_server/internal/repository/model"
	"github.com/Timofey335/jwt_server/internal/utils"
	descAuth "github.com/Timofey335/jwt_server/pkg/auth_v1"
)

func (s *serv) GetRefreshToken(ctx context.Context, req *descAuth.GetRefreshTokenRequest) (*descAuth.GetRefreshTokenResponse, error) {
	claims, err := utils.VerifyToken(req.GetRefreshToken(), []byte(refreshTokenSecretKey))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	user, err := s.userRepository.GetUserData(ctx, claims.Username)
	if err != nil {
		return nil, err
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

	return &descAuth.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
