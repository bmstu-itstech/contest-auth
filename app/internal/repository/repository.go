package repository

import (
	"context"

	"github.com/bmstu-itstech/contest-auth/internal/model"
)

type UserRepository interface {
	// CreateUser creates a new user in the database.
	CreateUser(ctx context.Context, params model.CreateUserParams) (resp model.CreateUserResponse, err error)

	// GetUserByEmail retrieves a user by their email.
	GetUserByEmail(ctx context.Context, params model.GetUserByEmailParams) (resp model.GetUserByEmailResponse, err error)
}
