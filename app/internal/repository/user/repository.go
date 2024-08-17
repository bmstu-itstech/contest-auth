package user

import (
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
