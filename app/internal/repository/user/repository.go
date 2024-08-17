package user

import (
	"context"

	"github.com/bmstu-itstech/contest-auth/internal/model"
	"github.com/bmstu-itstech/contest-auth/internal/repository"
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
	// TODO: Implement database logic for creating a user
	return model.CreateUserResponse{}, nil
}

// GetUserByEmail retrieves a user by their email.
func (r *userPGRepo) GetUserByEmail(
	ctx context.Context,
	params model.GetUserByEmailParams,
) (resp model.GetUserByEmailResponse, err error) {
	// TODO: Implement database logic for retrieving a user by email
	return model.GetUserByEmailResponse{}, nil
}
