package user

import (
	"context"

	"github.com/pkg/errors"

	"github.com/bmstu-itstech/contest-auth/internal/model"
	"github.com/bmstu-itstech/contest-auth/internal/repository"
	"github.com/bmstu-itstech/contest-auth/internal/repository/user/converter"
	modelRepo "github.com/bmstu-itstech/contest-auth/internal/repository/user/model"
	"github.com/bmstu-itstech/contest-auth/pkg/db"
)

type userPGRepo struct {
	db db.Client
}

// NewRepository creates a new instance of userPGRepo with the provided database connection.
func NewRepository(db db.Client) repository.UserRepository {
	return &userPGRepo{db: db}
}

// CreateUser creates a new user in the database.
func (r *userPGRepo) CreateUser(
	ctx context.Context,
	params model.CreateUserParams,
) (resp model.CreateUserResponse, err error) {
	paramsRepo := converter.ConvertCreateUserParamsFromServiceToRepo(params)

	var respRepo modelRepo.CreateUserResponse

	q := db.Query{
		Name:     "userPGRepo.CreateUser",
		QueryRaw: queryCreateUser,
	}

	err = r.db.DB().ScanOneContext(ctx, &respRepo, q, paramsRepo.Email, paramsRepo.Username, paramsRepo.PasswordHash)
	if err != nil {
		return resp, errors.Wrapf(
			err,
			"Cannot create user(email: %s)",
			paramsRepo.Email,
		)
	}

	return converter.ConvertCreateUserResponseFromRepoToService(respRepo), nil
}

// GetUserByEmail retrieves a user by their email.
func (r *userPGRepo) GetUserByEmail(
	ctx context.Context,
	params model.GetUserByEmailParams,
) (resp model.GetUserByEmailResponse, err error) {
	paramsRepo := converter.ConvertGetUserByEmailParamsFromServiceToRepo(params)

	var respRepo modelRepo.GetUserByEmailResponse

	q := db.Query{
		Name:     "userPGRepo.GetUserByEmail",
		QueryRaw: queryGetUserByEmail,
	}

	err = r.db.DB().ScanOneContext(ctx, &respRepo, q, paramsRepo.Email)
	if err != nil {
		return resp, errors.Wrapf(
			err,
			"Cannot create user(email: %s)",
			paramsRepo.Email,
		)
	}

	return converter.ConvertGetUserByEmailResponseFromRepoToService(respRepo), nil
}
