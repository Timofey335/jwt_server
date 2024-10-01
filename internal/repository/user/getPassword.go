package user

import (
	sq "github.com/Masterminds/squirrel"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	repoModel "github.com/Timofey335/jwt_server/internal/repository/model"
	descAuth "github.com/Timofey335/jwt_server/pkg/auth_v1"
)

func (r *repo) GetPassword(ctx context.Context, req *descAuth.LoginRequest) (*repoModel.UserRepoModel, error) {
	builderSelect := sq.Select(passwordColumn, roleColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{nameColumn: req.Username})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	// q := db.Query{
	// 	Name:     "user_repository.GetUserData",
	// 	QueryRaw: query,
	// }

	var password string
	var role int64
	// err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	err = r.db.QueryRow(ctx, query, args...).Scan(&password, &role)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &repoModel.UserRepoModel{
		Password: password,
		Role:     role,
	}, nil
}
