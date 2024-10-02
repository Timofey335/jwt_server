package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Timofey335/jwt_server/internal/repository/user"
	"github.com/Timofey335/jwt_server/internal/service"
	userService "github.com/Timofey335/jwt_server/internal/service/user"
	descAccess "github.com/Timofey335/jwt_server/pkg/access_v1"
	descAuth "github.com/Timofey335/jwt_server/pkg/auth_v1"
)

const (
	dbDSN      = "host=localhost port=54321 dbname=users user=user password=userspassword sslmode=disable"
	grpcPort   = 50051
	authPrefix = "Bearer "

	accessTokenSecretKey = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="

	accessTokenExpiration = 5 * time.Minute
)

type serverAuth struct {
	descAuth.UnimplementedAuthV1Server
	userService service.UserService
}

type serverAccess struct {
	descAccess.UnimplementedAccessV1Server
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userRepo := user.NewRepository(pool)
	userService := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	descAuth.RegisterAuthV1Server(s, &serverAuth{userService: userService})
	descAccess.RegisterAccessV1Server(s, &serverAccess{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *serverAuth) Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	refreshToken, err := s.userService.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(refreshToken.RefreshToken)

	return &descAuth.LoginResponse{
		RefreshToken: refreshToken.RefreshToken,
	}, nil
}

func (s *serverAuth) GetRefreshToken(ctx context.Context, req *descAuth.GetRefreshTokenRequest) (*descAuth.GetRefreshTokenResponse, error) {
	refreshToken, err := s.userService.GetRefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return &descAuth.GetRefreshTokenResponse{
		RefreshToken: refreshToken.RefreshToken,
	}, nil
}
